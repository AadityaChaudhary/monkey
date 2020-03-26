package parser

import (
	"monkey/token"
	"monkey/lexer"
	"monkey/ast"
)

type Parser struct {
	l 			*lexer.Lexer
	curToken	token.Token
	peekToken	token.Token

}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		statement := p.ParseStatement()

		if statement != nil {
			program.Statements = append(program.Statements,statement)
		}
		p.nextToken()
	}

	return program

}

func(p *Parser) ParseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.ParseLetStatement()
	default:
		return nil
	}

}

func(p *Parser) ParseLetStatement() *ast.LetStatement {
	statement := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	statement.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	for !p.currIs(token.EOF) {
		p.nextToken()
	}

	return statement

}

func(p *Parser) peekIs(tokenType token.TokenType) bool {
	return p.peekToken.Type == tokenType
}

func(p *Parser) currIs(tokenType token.TokenType) bool {
	return p.curToken.Type == tokenType
}

func(p *Parser) expectPeek(tokenType token.TokenType) bool {
	if p.peekToken.Type == tokenType {
		p.nextToken()
		return true
	} else {
		return false
	}
}
