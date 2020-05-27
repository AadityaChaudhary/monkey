package parser

import (
	"monkey/ast"
	"monkey/lexer"

	"testing"
)
func TestLetStatement(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 993322;`

	l := lexer.New(input)

	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("returned nil")
	}
	if len(program.Statements) != 3 {

		t.Fatalf("program should have 3 statements, got %d", len(program.Statements))
	}
	//tests := [] struct {
	//	expectedIdentifier string
	//}{
	//	{"x"},
	//	{"y"},
	//	{"foobar"},
	//}

	for _, statement := range program.Statements {
		returnStatement, ok := statement.(*ast.ReturnStatement)

		if !ok {
			t.Errorf("statement not a return statement, got %T", statement)
			continue
		}
		if returnStatement.TokenLiteral() != "return" {
			t.Errorf("returnStatement.TokenLiteral does not return 'return', got %q", returnStatement.TokenLiteral())
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
