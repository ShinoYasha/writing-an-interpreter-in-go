package repl

import (
	"bufio"
	"fmt"
	"interpreter/evaluator"
	"interpreter/lexer"
	"interpreter/parser"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		//for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		//	fmt.Printf("%+v\n", tok)
		//}
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}
		evaluated := evaluator.Eval(program)
		//io.WriteString(out, program.String())
		//io.WriteString(out, "\n")
		if evaluated != nil {
			_, _ = io.WriteString(out, evaluated.Inspect())
			_, _ = io.WriteString(out, "\n")
		}
	}
}
func printParseErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Woops!We ran into some errors here!\n")
	io.WriteString(out, " parse errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\t")
	}
}
