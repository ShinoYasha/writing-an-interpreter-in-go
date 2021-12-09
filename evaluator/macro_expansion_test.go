package evaluator

import (
	"interpreter/ast"
	"interpreter/lexer"
	"interpreter/object"
	"interpreter/parser"
	"testing"
)

func TestExpandMacro(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`
		let infixExpression = macro() { quote(1 + 2);};
		infixExpression();
		`,
			`(1 + 2)`,
		},
		{
			`
			let reverse = macro(a, b) { quote(unquote(b) - unquote(a));};
			reverse(2 + 2, 10 - 5);
			`,
			`(10 - 5) - (2 + 2)`,
		},
		{
			`
let unless = macro(condition, consequence, alternative) {
    quote(if (!(unquote(condition))) {
        unquote(consequence);
    } else {
        unquote(alternative);
}); };
unless(10 > 5, puts("not greater"), puts("greater"));
`,
			`if (!(10 > 5)) { puts("not greater") } else { puts("greater") }`,
		},
	}
	for _, tt := range tests {
		expected := testParseProgram(tt.expected)
		program := testParseProgram(tt.input)
		env := object.NewEnvironment()
		DefineMacros(program, env)
		expanded := ExpandMacro(program, env)
		if expected.String() != expanded.String() {
			t.Errorf("not equal. want=%q, got=%q", expected.String(), expanded.String())
		}

	}
}

func testParseProgram(input string) *ast.Program {
	l := lexer.New(input)
	p := parser.New(l)
	return p.ParseProgram()
}
