package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"testing"
)
func Test(t *testing.T) {

	input := `5;`

	l := lexer.New(input)

	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("returned nil")
	}
	if len(program.Statements) != 1 {

		t.Fatalf("program should have 1 statements, got %d", len(program.Statements))
	}



		statement, ok := program.Statements[0].(*ast.ExpressionStatement)

		if !ok {
			t.Errorf("statement not am identifier, got %T", statement)

		}
		integer,ok := statement.Expression.(*ast.IntegerLiteral)

		if !ok {
			t.Errorf("expression not an identifier, got %t",statement.Expression)
		}
		if integer.Value != 5 {
			t.Errorf("value not 5, got = %d", integer.Value)
		}
		if integer.TokenLiteral() != "5" {
			t.Errorf("token literal not foobar, got = %s", integer.TokenLiteral())
		}


}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integer, ok := il.(*ast.IntegerLiteral)

	if !ok {
		t.Errorf("not an integer literal")
	}
}

func checkParserErrors( t *testing.T, p *Parser) {
	errors  := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))

	for _,msg := range errors {
		t.Errorf("parser error: %q", msg)
	}

	t.FailNow()

}
