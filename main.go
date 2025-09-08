package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"pepper/lexer"
	"pepper/parser"
	"github.com/davecgh/go-spew/spew"
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

	spew.Dump(program)
}