package main

import (
	"flag"
	"fmt"
	"os"
	"pepper/compiler"
	"pepper/lexer"
	"pepper/parser"
	"pepper/runtime"
	"pepper/utils"
	"sync"
)

const version = "0.2.0"

func main() {
	var debug bool
	var verboseDebug bool
	var showVersion bool
	var dump string
	var exec string
	var human string

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: pepper [options] <file>\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}

	flag.BoolVar(&debug, "d", false, "enable debug mode")
	flag.BoolVar(&verboseDebug, "vd", false, "enable verbose debug mode (dumps VM instructions)")
	flag.BoolVar(&showVersion, "v", false, "show version information")
	flag.StringVar(&dump, "p", "", "dump bytecode to file")
	flag.StringVar(&exec, "e", "", "execute bytecode from file")
	flag.StringVar(&human, "h", "", "view dumped bytecode")
	flag.Parse()

	if verboseDebug {
		debug = true
	}

	if showVersion {
		fmt.Printf("Pepper v%s\n", version)
		return
	}

	if human != "" {
		data, err := os.ReadFile(human)
		if err != nil {
			panic(err)
		}

		instrs, err := utils.DecodeBytecode(data)
		if err != nil {
			panic(err)
		}

		fmt.Println("Instructions:")
		for i, instr := range instrs {
			fmt.Printf("%04d %s\n", i, runtime.ResolveVMInstruction(instr))
		}
		return
	}

	if exec != "" {
		data, err := os.ReadFile(exec)
		if err != nil {
			panic(err)
		}

		instrs, err := utils.DecodeBytecode(data)
		if err != nil {
			panic(err)
		}

		var wg sync.WaitGroup
		vm := runtime.NewVM(instrs, &wg)
		vm.Run(debug, verboseDebug)
		wg.Wait()
		return
	}

	filePath := flag.Arg(0)
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	l := lexer.New(string(data))
	p := parser.New(l)

	program := p.ParseProgram()
	comp := compiler.NewCompiler().Compile(program, false)

	if dump != "" {
		encoded, err := utils.EncodeBytecode(comp)
		if err != nil {
			panic(err)
		}
		err = os.WriteFile(dump, encoded, 0644)
		if err != nil {
			panic(err)
		}
		return
	}

	var wg sync.WaitGroup
	vm := runtime.NewVM(comp, &wg)
	vm.Run(debug, verboseDebug)

	if debug {
		fmt.Println("Stack:")
		runtime.DumpOperandStack(vm)
		fmt.Println("Memory:")
		runtime.DumpMemory(vm)
	}

	wg.Wait()
}
