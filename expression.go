package main

type Expr interface {
	Accept(v Visitor) any
}

type Visitor interface {
	VisitBinary(expr *Binary) any
	VisitGrouping(expr *Grouping) any
	VisitLiteral(expr *Literal) any
	VisitUnary(expr *Unary) any
}

type Binary struct {
	Left, Right Expr
	Operator    Token
}

func (b *Binary) Accept(v Visitor) any {
	return v.VisitBinary(b)
}

type Grouping struct {
	Expression Expr
}

func (g *Grouping) Accept(v Visitor) any {
	return v.VisitGrouping(g)
}

type Literal struct {
	Value any
}

func (l *Literal) Accept(v Visitor) any {
	return v.VisitLiteral(l)
}

type Unary struct {
	Operator Token
	Right    Expr
}

func (u *Unary) Accept(v Visitor) any {
	return v.VisitUnary(u)
}
