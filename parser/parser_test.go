package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"testing"
)
func Test(t *testing.T) {
	tests := []struct {
		input 		string
		leftVal		int64
		operator 	string
		rightVal 	int64
	}{
		{"5 + 5;", 5, "+",5},
		{"5 - 5;", 5, "-",5},
		{"5 * 5;", 5, "*",5},
		{"5 / 5;", 5, "/",5},
		{"5 > 5;", 5, ">",5},
		{"5 < 5;", 5, "<",5},
		{"5 == 5;", 5, "==",5},
		{"5 != 5;", 5, "!=",5},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)

		program := p.ParseProgram()
		checkParserErrors(t,p)

		if len(program.Statements) != 1 {
			t.Fatalf("got %d statements", len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("stmt isnt an ast.ExpressionStatement, got= %T", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("stmt is not an infixexpression, got = %T", stmt.Expression)
		}

		if !testIntegerLiteral(t,exp.Left, tt.leftVal) {
			return
		}

		if exp.Operator != tt.operator {
			t.Fatalf("exp operator is not %s, got= %s", tt.operator, exp.Operator)
		}

		if !testIntegerLiteral(t,exp.Right, tt.rightVal) {
			return
		}
	}

}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integer, ok := il.(*ast.IntegerLiteral)

	if !ok {
		t.Errorf("not an integer literal")
		return false
	}

	if integer.Value != value {
		t.Errorf("integer value not %d, got =%d", value, integer.Value)
		return false
	}

	if integer.TokenLiteral() != fmt.Sprintf("%d",value) {
		t.Errorf("integer token literal not %d, got %s", value, integer.TokenLiteral())
		return false
	}

	return true
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
