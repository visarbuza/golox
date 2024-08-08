package main

type Expr interface {
	Accept(v ExprVisitor) any
}

type ExprVisitor interface {
	VisitBinaryExpr(expr *Binary) any
	VisitGroupingExpr(expr *Grouping) any
	VisitLiteralExpr(expr *Literal) any
	VisitUnaryExpr(expr *Unary) any
	VisitVariableExpr(expr *Variable) any
	VisitAssignExpr(expr *Assign) any
	VisitLogicalExpr(expr *Logical) any
}

type Binary struct {
	Left, Right Expr
	Operator    Token
}

func (b *Binary) Accept(v ExprVisitor) any {
	return v.VisitBinaryExpr(b)
}

type Grouping struct {
	Expression Expr
}

func (g *Grouping) Accept(v ExprVisitor) any {
	return v.VisitGroupingExpr(g)
}

type Literal struct {
	Value any
}

func (l *Literal) Accept(v ExprVisitor) any {
	return v.VisitLiteralExpr(l)
}

type Unary struct {
	Operator Token
	Right    Expr
}

func (u *Unary) Accept(v ExprVisitor) any {
	return v.VisitUnaryExpr(u)
}

type Variable struct {
	Name Token
}

func (v *Variable) Accept(ev ExprVisitor) any {
	return ev.VisitVariableExpr(v)
}

type Assign struct {
	Name  Token
	Value Expr
}

func (a *Assign) Accept(ev ExprVisitor) any {
	return ev.VisitAssignExpr(a)
}

type Logical struct {
	Left     Expr
	Operator Token
	Right    Expr
}

func (l *Logical) Accept(v ExprVisitor) any {
	return v.VisitLogicalExpr(l)
}
