package main

import (
	"bufio"
	"fmt"
	"os"
	"pepper/compiler"
	"pepper/lexer"
	"pepper/parser"
	"pepper/runtime"
	"strings"
	"sync"
)

func Prompt() {
	fmt.Println("Pepper REPL")
	fmt.Println("Type .quit to exit")
	fmt.Println("Type .run to run the program")
	fmt.Println("Type .list to list the program")
	fmt.Println("Type .clear to clear the program")

	program := ""
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			return
		}

		line = strings.TrimSpace(line)

		switch line {
		case ".quit":
			return
		case ".run":
			l := lexer.New(program)
			p := parser.New(l)

			parsedProgram := p.ParseProgram()
			if len(p.Errors()) != 0 {
				for _, msg := range p.Errors() {
					fmt.Fprintln(os.Stderr, msg)
				}
				continue
			}

			comp := compiler.NewCompiler()
			instr := comp.Compile(parsedProgram)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Woops! Compilation failed:", err)
				continue
			}

			var wg sync.WaitGroup
			vm := runtime.NewVM(instr, &wg)
			vm.Run(false, false)

		case ".list":
			fmt.Println(program)
		case ".clear":
			program = ""
		default:
			program += line + "\n"
		}
	}
}
