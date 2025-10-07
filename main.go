package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"pepper/compiler"
	"pepper/lexer"
	"pepper/parser"
	"pepper/runtime"
	"sync"
)

func main() {
	if len(os.Args) == 3 && os.Args[2] != "-d" {
		fmt.Println("Usage: pepper <file> [-d for debug]")
		return
	}

	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	l := lexer.New(string(data))
	p := parser.New(l)

	program := p.ParseProgram()
	comp := compiler.NewCompiler().Compile(program)

	if len(os.Args) == 3 && os.Args[2] == "-d" {
		fmt.Println("Instructions:")
		for i, instr := range comp {
			fmt.Printf("%04d %s\n", i, runtime.ResolveVMInstruction(instr))
		}
	}

	var wg sync.WaitGroup
	vm := runtime.NewVM(comp, &wg)
	if len(os.Args) == 3 && os.Args[2] == "-d" {
		vm.Run(true)
	} else {
		vm.Run(false)
	}

	if len(os.Args) == 3 && os.Args[2] == "-d" {
		fmt.Println("Stack:")
		runtime.DumpOperandStack(vm)
		fmt.Println("Memory:")
		runtime.DumpMemory(vm)
	}

	wg.Wait()
}
