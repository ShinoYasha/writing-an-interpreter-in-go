package evaluator

import (
	"interpreter/object"
	"testing"
)

func TestQuoteUnquote(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`quote(unquote(5))`,
			`5`,
		},
		{
			`quote(unquote(4 + 4))`,
			`8`,
		},
		{
			`quote(8 + unquote(4 + 4))`,
			`(8 + 8)`,
		},
		{
			`quote(unquote(4 + 4) + 4)`,
			`(8 + 4)`,
		},
		{
			`let foobar = 8;
quote(foobar)`,
			`foobar`,
		},
		{
			`let foobar = 8;
			quote(unquote(foobar))`,
			`8`,
		},
		{
			`quote(unquote(true))`,
			`true`,
		},
		{
			`quote(unquote(true == false))`,
			`false`,
		},
		{
			`quote(unquote(quote(4 + 4)))`,
			`(4 + 4)`,
		},
		{
			`let quoteInfixExpression = quote(4 + 4);
			quote(unquote(4 + 4) + unquote(quoteInfixExpression))`,
			`(8 + (4 + 4))`,
		},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		quote, ok := evaluated.(*object.Quote)
		if !ok {
			t.Fatalf("expected *object.Quote.got=%T(%+v)",
				evaluated, evaluated)
		}
		if quote.Node == nil {
			t.Fatalf("quote.Node is nil")
		}
		if quote.Node.String() != tt.expected {
			t.Errorf("not equal. got=%q, want=%q",
				quote.Node.String(), tt.expected)
		}
	}
}
