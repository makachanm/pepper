package main

import (
	"flag"
	"fmt"
	"os"
	"pepper/compiler"
	"pepper/lexer"
	"pepper/parser"
	"pepper/runtime"
	"sync"
)

const version = "0.1.1"

func main() {
	var debug bool
	var showVersion bool

	flag.BoolVar(&debug, "d", false, "enable debug mode")
	flag.BoolVar(&showVersion, "v", false, "show version information")
	flag.Parse()

	if showVersion {
		fmt.Printf("Pepper v%s\n", version)
		return
	}

	if len(flag.Args()) != 1 {
		fmt.Println("Usage: pepper [-d] [-v] <file>")
		os.Exit(1)
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

	if debug {
		fmt.Println("Instructions:")
		for i, instr := range comp {
			fmt.Printf("%04d %s\n", i, runtime.ResolveVMInstruction(instr))
		}
	}

	var wg sync.WaitGroup
	vm := runtime.NewVM(comp, &wg)
	vm.Run(debug)

	if debug {
		fmt.Println("Stack:")
		runtime.DumpOperandStack(vm)
		fmt.Println("Memory:")
		runtime.DumpMemory(vm)
	}

	wg.Wait()
}
