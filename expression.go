package main

type Expr interface {
	Accept(v ExprVisitor) any
}

type ExprVisitor interface {
	VisitBinary(expr *Binary) any
	VisitGrouping(expr *Grouping) any
	VisitLiteral(expr *Literal) any
	VisitUnary(expr *Unary) any
	VisitVariable(expr *Variable) any
}

type Binary struct {
	Left, Right Expr
	Operator    Token
}

func (b *Binary) Accept(v ExprVisitor) any {
	return v.VisitBinary(b)
}

type Grouping struct {
	Expression Expr
}

func (g *Grouping) Accept(v ExprVisitor) any {
	return v.VisitGrouping(g)
}

type Literal struct {
	Value any
}

func (l *Literal) Accept(v ExprVisitor) any {
	return v.VisitLiteral(l)
}

type Unary struct {
	Operator Token
	Right    Expr
}

func (u *Unary) Accept(v ExprVisitor) any {
	return v.VisitUnary(u)
}

type Variable struct {
	Name Token
}

func (v *Variable) Accept(ev ExprVisitor) any {
	return ev.VisitVariable(v)
}
