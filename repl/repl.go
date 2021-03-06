package repl

import (
	"bufio"
	"fmt"
	"interpreter/compiler"
	"interpreter/lexer"
	"interpreter/parser"
	"interpreter/vm"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	//env := object.NewEnvironment()
	//macroEnv := object.NewEnvironment()
	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}
		//evaluator.DefineMacros(program, macroEnv)
		//expanded := evaluator.ExpandMacro(program, macroEnv)
		//evaluated := evaluator.Eval(expanded, env)
		//if evaluated != nil {
		//	_, _ = io.WriteString(out, evaluated.Inspect())
		//	_, _ = io.WriteString(out, "\n")
		//}
		comp := compiler.New()
		err := comp.Compile(program)
		if err != nil {
			fmt.Fprintf(out, "Woops! Compilation failed:\n %s\n", err)
			continue
		}
		machine := vm.New(comp.ByteCode())
		err = machine.Run()
		if err != nil {
			fmt.Fprintf(out, "Woops! Executing bytecode failed:\n %s\n", err)
		}
		stackTop := machine.LastPoppedStackElem()
		io.WriteString(out, stackTop.Inspect())
		io.WriteString(out, "\n")
	}
}
func printParseErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Woops!We ran into some errors here!\n")
	io.WriteString(out, " parse errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\t")
	}
}
