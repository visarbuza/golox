package main

import (
	"fmt"
)

type Interpreter struct {
}

func (i *Interpreter) Interpret(expr Expr) {
	defer func() {
		if r := recover(); r != nil {
			val, _ := r.(RuntimeError)
			runtimeError(val)
		}
	}()
	val := i.evaluate(expr)
	fmt.Println(i.stringify(val))
}

func (i *Interpreter) VisitBinary(expr *Binary) any {
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

func (i *Interpreter) VisitGrouping(expr *Grouping) any {
	return i.evaluate(expr.Expression)
}

func (i *Interpreter) VisitLiteral(expr *Literal) any {
	return expr.Value
}

func (i *Interpreter) VisitUnary(expr *Unary) any {
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
	return true
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
