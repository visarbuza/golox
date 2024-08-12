package main

type Stmt interface {
	Accept(v StmtVisitor) any
}

type StmtVisitor interface {
	VisitPrintStmt(stmt *PrintStmt) any
	VisitExpressionStmt(stmt *ExpressionStmt) any
	VisitVarStmt(stmt *VarStmt) any
	VisitBlockStmt(stmt *Block) any
	VisitIfStmt(stmt *IfStmt) any
	VisitWhileStmt(stmt *WhileStmt) any
	VisitBreakStmt(stmt *BreakStmt) any
	VisitFunctionStmt(stmt *FunctionStmt) any
	VisitReturnStmt(stmt *ReturnStmt) any
}

type Block struct {
	Statements []Stmt
}

func (b *Block) Accept(v StmtVisitor) any {
	return v.VisitBlockStmt(b)
}

type PrintStmt struct {
	Expression Expr
}

func (p *PrintStmt) Accept(v StmtVisitor) any {
	return v.VisitPrintStmt(p)
}

type ExpressionStmt struct {
	Expression Expr
}

func (e *ExpressionStmt) Accept(v StmtVisitor) any {
	return v.VisitExpressionStmt(e)
}

type VarStmt struct {
	Name        Token
	Initializer Expr
}

func (e *VarStmt) Accept(v StmtVisitor) any {
	return v.VisitVarStmt(e)
}

type IfStmt struct {
	Condition  Expr
	ThenBranch Stmt
	ElseBranch Stmt
}

func (i *IfStmt) Accept(v StmtVisitor) any {
	return v.VisitIfStmt(i)
}

type WhileStmt struct {
	Condition Expr
	Body      Stmt
}

func (w *WhileStmt) Accept(v StmtVisitor) any {
	return v.VisitWhileStmt(w)
}

type BreakStmt struct {
	Keyword Token
}

func (b *BreakStmt) Accept(v StmtVisitor) any {
	return v.VisitBreakStmt(b)
}

type FunctionStmt struct {
	Name   Token
	Params []Token
	Body   []Stmt
}

func (f *FunctionStmt) Accept(v StmtVisitor) any {
	return v.VisitFunctionStmt(f)
}

type ReturnStmt struct {
	Keyword Token
	Value   Expr
}

func (r *ReturnStmt) Accept(v StmtVisitor) any {
	return v.VisitReturnStmt(r)
}
