package main

import (
	"bytes"
	"fmt"
)

type AstPrinter struct {
}

func (ap *AstPrinter) Print(expr Expr) string {
	return expr.Accept(ap).(string)
}

func (ap *AstPrinter) VisitBinary(expr *Binary) any {
	return ap.parenthesize(expr.Operator.Lexeme, expr.Left, expr.Right)
}

func (ap *AstPrinter) VisitGrouping(expr *Grouping) any {
	return ap.parenthesize("group", expr.Expression)
}

func (ap *AstPrinter) VisitLiteral(expr *Literal) any {
	if expr.Value == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", expr.Value)
}

func (ap *AstPrinter) VisitUnary(expr *Unary) any {
	return ap.parenthesize(expr.Operator.Lexeme, expr.Right)
}

func (ap *AstPrinter) parenthesize(name string, exprs ...Expr) string {
	var buffer bytes.Buffer
	buffer.WriteString("(")
	buffer.WriteString(name)
	for _, expr := range exprs {
		buffer.WriteString(" ")
		buffer.WriteString(expr.Accept(ap).(string))
	}
	buffer.WriteString(")")
	return buffer.String()
}

type RPNPrinter struct {
}

func (ap *RPNPrinter) Print(expr Expr) string {
	return expr.Accept(ap).(string)
}

func (ap *RPNPrinter) VisitBinary(expr *Binary) any {
	return ap.rpn(expr.Operator.Lexeme, expr.Left, expr.Right)
}

func (ap *RPNPrinter) VisitGrouping(expr *Grouping) any {
	return ap.rpn("", expr.Expression)
}

func (ap *RPNPrinter) VisitLiteral(expr *Literal) any {
	if expr.Value == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", expr.Value)
}

func (ap *RPNPrinter) VisitUnary(expr *Unary) any {
	return ap.rpn(expr.Operator.Lexeme, expr.Right)
}

func (ap *RPNPrinter) rpn(name string, exprs ...Expr) string {
	var buffer bytes.Buffer
	for _, expr := range exprs {
		buffer.WriteString(expr.Accept(ap).(string))
		buffer.WriteString(" ")
	}
	buffer.WriteString(name)
	return buffer.String()
}
