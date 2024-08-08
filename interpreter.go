package main

import (
	"fmt"
)

type Interpreter struct {
	env *Environment
}

func (i *Interpreter) Interpret(statements []Stmt) {
	defer func() {
		if r := recover(); r != nil {
			val, _ := r.(RuntimeError)
			runtimeError(val)
		}
	}()
	for _, stmt := range statements {
		i.execute(stmt)
	}
}

func (i *Interpreter) VisitPrintStmt(stmt *PrintStmt) any {
	val := i.evaluate(stmt.Expression)
	fmt.Println(i.stringify(val))
	return nil
}

func (i *Interpreter) VisitExpressionStmt(stmt *ExpressionStmt) any {
	i.evaluate(stmt.Expression)
	return nil
}

func (i *Interpreter) VisitVarStmt(stmt *VarStmt) any {
	var value any = nil
	if stmt.Initializer != nil {
		value = i.evaluate(stmt.Initializer)
	}
	i.env.Define(stmt.Name.Lexeme, value)
	return nil
}

func (i *Interpreter) VisitBlockStmt(stmt *Block) any {
	i.executeBlock(stmt.Statements, NewEnvironment(i.env))
	return nil
}

func (i *Interpreter) VisitIfStmt(stmt *IfStmt) any {
	if i.isTruthy(i.evaluate(stmt.Condition)) {
		i.execute(stmt.ThenBranch)
	} else if stmt.ElseBranch != nil {
		i.execute(stmt.ElseBranch)
	}
	return nil
}

func (i *Interpreter) VisitWhileStmt(stmt *WhileStmt) any {
	for i.isTruthy(i.evaluate(stmt.Condition)) {
		i.execute(stmt.Body)
	}
	return nil
}

func (i *Interpreter) VisitBinaryExpr(expr *Binary) any {
	left := i.evaluate(expr.Left)
	right := i.evaluate(expr.Right)

	switch expr.Operator.Type {
	case MINUS:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) - right.(float64)
	case SLASH:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) / right.(float64)
	case STAR:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) * right.(float64)
	case PLUS:
		if left, ok := left.(float64); ok {
			if right, ok := right.(float64); ok {
				return left + right
			}
		}
		if left, ok := left.(string); ok {
			if right, ok := right.(string); ok {
				return left + right
			}
		}
		panic(RuntimeError{expr.Operator, "Operands must be two numbers or two strings"})
	case GREATER:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) > right.(float64)
	case GREATER_EQUAL:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) >= right.(float64)
	case LESS:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) < right.(float64)
	case LESS_EQUAL:
		i.checkNumberOperands(expr.Operator, left, right)
		return left.(float64) <= right.(float64)
	case BANG_EQUAL:
		return !i.isEqual(left, right)
	case EQUAL_EQUAL:
		return i.isEqual(left, right)
	}
	return nil
}

func (i *Interpreter) VisitGroupingExpr(expr *Grouping) any {
	return i.evaluate(expr.Expression)
}

func (i *Interpreter) VisitLiteralExpr(expr *Literal) any {
	return expr.Value
}

func (i *Interpreter) VisitUnaryExpr(expr *Unary) any {
	right := i.evaluate(expr.Right)

	switch expr.Operator.Type {
	case MINUS:
		i.checkNumberOperand(expr.Operator, right)
		return -right.(float64)
	case BANG:
		return !i.isTruthy(right)
	}

	return nil
}

func (i *Interpreter) VisitVariableExpr(expr *Variable) any {
	val, err := i.env.Get(expr.Name)
	if err != nil {
		panic(err)
	}
	return val
}

func (i *Interpreter) VisitAssignExpr(expr *Assign) any {
	value := i.evaluate(expr.Value)
	i.env.Assign(expr.Name, value)
	return value
}

func (i *Interpreter) VisitLogicalExpr(expr *Logical) any {
	left := i.evaluate(expr.Left)

	if expr.Operator.Type == OR {
		if i.isTruthy(left) {
			return left
		}
	} else {
		if !i.isTruthy(left) {
			return left
		}
	}
	return i.evaluate(expr.Right)
}

func (i *Interpreter) execute(stmt Stmt) {
	stmt.Accept(i)
}

func (i *Interpreter) executeBlock(statements []Stmt, env *Environment) {
	prevEnv := i.env
	i.env = env
	defer func() {
		i.env = prevEnv
	}()
	for _, stmt := range statements {
		i.execute(stmt)
	}
}

func (i *Interpreter) evaluate(expr Expr) any {
	return expr.Accept(i)
}

func (i *Interpreter) isTruthy(value any) bool {
	if value == nil {
		return false
	}
	ok, val := value.(bool)
	if ok {
		return val
	}
	return false
}

func (i *Interpreter) isEqual(a, b any) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil {
		return false
	}
	return a == b
}

func (i *Interpreter) stringify(value any) string {
	if value == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", value)
}

func (i *Interpreter) checkNumberOperand(operator Token, operand any) {
	if _, ok := operand.(float64); ok {
		return
	}
	panic(RuntimeError{operator, "Operand must be a number"})
}

func (i *Interpreter) checkNumberOperands(operator Token, left, right any) {
	if _, ok := left.(float64); ok {
		if _, ok := right.(float64); ok {
			return
		}
	}
	panic(RuntimeError{operator, "Operands must be numbers"})
}

type RuntimeError struct {
	Token   Token
	Message string
}

func (r RuntimeError) Error() string {
	return r.Message
}
