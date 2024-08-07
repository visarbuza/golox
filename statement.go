package main

type Stmt interface {
	Accept(v StmtVisitor) any
}

type StmtVisitor interface {
	VisitPrintStmt(stmt *PrintStmt) any
	VisitExpressionStmt(stmt *ExpressionStmt) any
	VisitVarStmt(stmt *VarStmt) any
	VisitBlockStmt(stmt *Block) any
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
