package main

import "strings"

type Expression interface {
	interpret(context string) bool
}

type TerminalExpression struct {
	data string
}

func NewTerminalExpression(data string) TerminalExpression {
	return TerminalExpression{data: data}
}
func (tme TerminalExpression) interpret(context string) bool {
	if strings.Contains(context, tme.data) {
		return true
	}
	return false
}

type OrExpression struct {
	expr1 Expression
	expr2 Expression
}

func NewOrExpression(expr1, expr2 Expression) OrExpression {
	return OrExpression{expr1: expr1, expr2: expr2}
}
func (oep OrExpression) interpret(context string) bool {
return oep.expr1.interpret(context)||oep.expr2.interpret(context)
}

type AndExpression struct {
	expr1 Expression
	expr2 Expression
}

func NewAndExpression(expr1, expr2 Expression) AndExpression {
	return AndExpression{expr1: expr1, expr2: expr2}
}
func (oep AndExpression) interpret(context string) bool {
	return oep.expr1.interpret(context)&&oep.expr2.interpret(context)
}
