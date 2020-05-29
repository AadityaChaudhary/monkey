package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
	"strconv"
)

type Parser struct {
	l 			*lexer.Lexer
	curToken	token.Token
	peekToken	token.Token
	errors 		[]string

	prefixParseFns	map[token.TokenType]prefixParseFn
	infixParseFn	map[token.TokenType]infixParseFn
}

const (
	_ int = iota
	LOWEST
	EQUALS 		// ==
	LESSGREATER //> or <
	SUM 		//+
	PRODUCT 	//*
	PREFIX		//-X or !X
	CALL		//myFunction(X)
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn func(ast.Expression) ast.Expression
)

func(p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func(p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFn[tokenType] = fn
}

func(p *Parser) Errors() []string {
	return p.errors
}

func(p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)

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
	case token.RETURN:
		return p.ParseReturnStatement()
	default:
		return  p.parseExpressionStatement()
	}

}

func(p * Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func(p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)

	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value

	return lit
}

func(p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		return nil
	}
	leftExp := prefix()

	return leftExp
}

func(p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func(p *Parser) ParseReturnStatement() *ast.ReturnStatement {
	statement := &ast.ReturnStatement{Token: p.curToken}

	for !p.currIs(token.SEMICOLON) {
		p.nextToken()
	}

	return statement
}

func(p *Parser) ParseLetStatement() *ast.LetStatement {
	statement := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	//this is just the name
	statement.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	//ignoring the expression since we dont know how to parse that yet
	for !p.currIs(token.SEMICOLON) {
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
		p.peekError(tokenType)
		return false
	}
}
