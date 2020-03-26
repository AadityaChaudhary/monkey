package parser

import (
	"monkey/ast"
	"monkey/lexer"

	"testing"
)
func TestLetStatement(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;	`

	l := lexer.New(input)

	p := New(l)

	program := p.ParseProgram()

	if program == nil {
		t.Fatalf("returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program should have 3 statements, go %d", len(program.Statements))
	}
	tests := [] struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]

		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}

}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("token literal not 'let', got=%q", s.TokenLiteral())
		return false;
	}

	letStmt, ok := s.(*ast.LetStatement)

	if !ok {
		t.Errorf("s not a let statement, got %T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letstmt. name . value not '%s' got=%s",name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letstmt. name . TOKEN LITERAL not '%s' got=%s",name, letStmt.Name)
		return false
	}
	return true


}
