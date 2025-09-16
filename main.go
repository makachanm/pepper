package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"pepper/compiler"
	"pepper/lexer"
	"pepper/parser"
	"pepper/runtime"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: pepper <file>")
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

	vm := runtime.NewVM(comp)
	vm.Run()
}
