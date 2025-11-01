package parser

import (
	"pepper/lexer"
)

// The base Node interface
type Node interface {
	TokenLiteral() string
	GetToken() lexer.Token
}

// All statement nodes implement this
type Statement interface {
	Node
	statementNode()
}

// All expression nodes implement this
type Expression interface {
	Node
	expressionNode()
}

// The root node of every AST our parser produces.
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) GetToken() lexer.Token {
	if len(p.Statements) > 0 {
		return p.Statements[0].GetToken()
	}
	return lexer.Token{}
}

// Statements
type LetStatement struct {
	Token lexer.Token // the lexer.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()        {}
func (ls *LetStatement) TokenLiteral() string  { return ls.Token.Literal }
func (ls *LetStatement) GetToken() lexer.Token { return ls.Token }

type DimStatement struct {
	Token       lexer.Token // the lexer.DIM token
	Name        *Identifier
	AssignToken lexer.Token // the = token
	Value       Expression
}

func (ds *DimStatement) statementNode()        {}
func (ds *DimStatement) TokenLiteral() string  { return ds.Token.Literal }
func (ds *DimStatement) GetToken() lexer.Token { return ds.Token }

type ReturnStatement struct {
	Token       lexer.Token // the 'return' token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()        {}
func (rs *ReturnStatement) TokenLiteral() string  { return rs.Token.Literal }
func (rs *ReturnStatement) GetToken() lexer.Token { return rs.Token }

type ExpressionStatement struct {
	Token      lexer.Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode()        {}
func (es *ExpressionStatement) TokenLiteral() string  { return es.Token.Literal }
func (es *ExpressionStatement) GetToken() lexer.Token { return es.Token }

// Expressions
type Identifier struct {
	Token lexer.Token // the lexer.IDENT token
	Value string
}

func (i *Identifier) expressionNode()       {}
func (i *Identifier) TokenLiteral() string  { return i.Token.Literal }
func (i *Identifier) GetToken() lexer.Token { return i.Token }

type Boolean struct {
	Token lexer.Token
	Value bool
}

func (b *Boolean) expressionNode()       {}
func (b *Boolean) TokenLiteral() string  { return b.Token.Literal }
func (b *Boolean) GetToken() lexer.Token { return b.Token }

type IntegerLiteral struct {
	Token lexer.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()       {}
func (il *IntegerLiteral) TokenLiteral() string  { return il.Token.Literal }
func (il *IntegerLiteral) GetToken() lexer.Token { return il.Token }

type RealLiteral struct {
	Token lexer.Token
	Value float64
}

func (rl *RealLiteral) expressionNode()       {}
func (rl *RealLiteral) TokenLiteral() string  { return rl.Token.Literal }
func (rl *RealLiteral) GetToken() lexer.Token { return rl.Token }

type StringLiteral struct {
	Token lexer.Token
	Value string
}

func (sl *StringLiteral) expressionNode()       {}
func (sl *StringLiteral) TokenLiteral() string  { return sl.Token.Literal }
func (sl *StringLiteral) GetToken() lexer.Token { return sl.Token }

type NilLiteral struct {
	Token lexer.Token
}

func (nl *NilLiteral) expressionNode()       {}
func (nl *NilLiteral) TokenLiteral() string  { return nl.Token.Literal }
func (nl *NilLiteral) GetToken() lexer.Token { return nl.Token }

type FunctionLiteral struct {
	Token      lexer.Token // The 'func' token
	Name       *Identifier
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode()       {}
func (fl *FunctionLiteral) TokenLiteral() string  { return fl.Token.Literal }
func (fl *FunctionLiteral) GetToken() lexer.Token { return fl.Token }

type BlockStatement struct {
	Token      lexer.Token // the { token
	Statements []Statement
}

func (bs *BlockStatement) statementNode()        {}
func (bs *BlockStatement) TokenLiteral() string  { return bs.Token.Literal }
func (bs *BlockStatement) GetToken() lexer.Token { return bs.Token }

type BlockExpression struct {
	Token      lexer.Token // the { token
	Statements []Statement
}

func (be *BlockExpression) expressionNode()       {}
func (be *BlockExpression) TokenLiteral() string  { return be.Token.Literal }
func (be *BlockExpression) GetToken() lexer.Token { return be.Token }

type IfExpression struct {
	Token       lexer.Token // The 'if' token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode()       {}
func (ie *IfExpression) TokenLiteral() string  { return ie.Token.Literal }
func (ie *IfExpression) GetToken() lexer.Token { return ie.Token }

type CallExpression struct {
	Token     lexer.Token // The '(' token
	Function  Expression  // Identifier or FunctionLiteral
	Arguments []Expression
}

func (ce *CallExpression) expressionNode()       {}
func (ce *CallExpression) TokenLiteral() string  { return ce.Token.Literal }
func (ce *CallExpression) GetToken() lexer.Token { return ce.Token }

type PrefixExpression struct {
	Token    lexer.Token // The prefix token, e.g. !
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode()       {}
func (pe *PrefixExpression) TokenLiteral() string  { return pe.Token.Literal }
func (pe *PrefixExpression) GetToken() lexer.Token { return pe.Token }

type InfixExpression struct {
	Token    lexer.Token // The operator token, e.g. +
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode()       {}
func (ie *InfixExpression) TokenLiteral() string  { return ie.Token.Literal }
func (ie *InfixExpression) GetToken() lexer.Token { return ie.Token }

type AssignmentExpression struct {
	Token lexer.Token // The '=' token
	Left  Expression
	Value Expression
}

func (ae *AssignmentExpression) expressionNode()       {}
func (ae *AssignmentExpression) TokenLiteral() string  { return ae.Token.Literal }
func (ae *AssignmentExpression) GetToken() lexer.Token { return ae.Token }

type RepeatStatement struct {
	Token lexer.Token // The 'repeat' token
	Count Expression
	Body  *BlockStatement
}

func (rs *RepeatStatement) statementNode()        {}
func (rs *RepeatStatement) TokenLiteral() string  { return rs.Token.Literal }
func (rs *RepeatStatement) GetToken() lexer.Token { return rs.Token }

type LoopStatement struct {
	Token     lexer.Token // The 'loop' token
	Condition Expression
	Body      *BlockStatement
}

func (ls *LoopStatement) statementNode()        {}
func (ls *LoopStatement) TokenLiteral() string  { return ls.Token.Literal }
func (ls *LoopStatement) GetToken() lexer.Token { return ls.Token }

type BreakStatement struct {
	Token lexer.Token // The 'break' token
}

func (bs *BreakStatement) statementNode()        {}
func (bs *BreakStatement) TokenLiteral() string  { return bs.Token.Literal }
func (bs *BreakStatement) GetToken() lexer.Token { return bs.Token }

type ContinueStatement struct {
	Token lexer.Token // The 'continue' token
}

func (cs *ContinueStatement) statementNode()        {}
func (cs *ContinueStatement) TokenLiteral() string  { return cs.Token.Literal }
func (cs *ContinueStatement) GetToken() lexer.Token { return cs.Token }

type PackLiteral struct {
	Token lexer.Token // the '[' token
	Pairs map[Expression]Expression
}

func (pl *PackLiteral) expressionNode()       {}
func (pl *PackLiteral) TokenLiteral() string  { return pl.Token.Literal }
func (pl *PackLiteral) GetToken() lexer.Token { return pl.Token }

type IndexExpression struct {
	Token lexer.Token // the [ token
	Left  Expression
	Index Expression
}

func (ie *IndexExpression) expressionNode()       {}
func (ie *IndexExpression) TokenLiteral() string  { return ie.Token.Literal }
func (ie *IndexExpression) GetToken() lexer.Token { return ie.Token }

type MemberAccessExpression struct {
	Token  lexer.Token // The '->' token
	Object Expression
	Member *Identifier
}

func (mae *MemberAccessExpression) expressionNode()       {}
func (mae *MemberAccessExpression) TokenLiteral() string  { return mae.Token.Literal }
func (mae *MemberAccessExpression) GetToken() lexer.Token { return mae.Token }

type IncludeStatement struct {
	Token    lexer.Token // The 'include' token
	Filename string
}

func (is *IncludeStatement) statementNode()        {}
func (is *IncludeStatement) TokenLiteral() string  { return is.Token.Literal }
func (is *IncludeStatement) GetToken() lexer.Token { return is.Token }

type MatchExpression struct {
	Token      lexer.Token // The 'match' token
	Expression Expression
	Cases      []*MatchCase
}

func (me *MatchExpression) statementNode()        {}
func (me *MatchExpression) TokenLiteral() string  { return me.Token.Literal }
func (me *MatchExpression) GetToken() lexer.Token { return me.Token }

type MatchCase struct {
	Token   lexer.Token // The 'case' or 'default' token
	Pattern Expression
	Body    *BlockStatement
}
