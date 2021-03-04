package main

import (
	"fmt"

	"./simple_math"
)

type MathExpr = string

const (
	AddExp = MathExpr("add")
	SubExp = MathExpr("sub")
	MulExp = MathExpr("mul")
)

func main() {
	value := mathExpression(AddExp) // value is name of the func returned from mathExpression
	fmt.Println(value(5, 3))

	fmt.Println(double(7, 2, mathExpression(SubExp)))
}

// function takes string as param and return function
func mathExpression(expr MathExpr) func(f1, f2 float64) float64 {
	switch expr {
	case AddExp:
		return simple_math.Add // returning add func, not calling it. i.e., why not parentheses
	case SubExp:
		return simple_math.Subtract // returning subtract func, not calling it. i.e., why not parentheses
	case MulExp:
		return simple_math.Multiply // returning multiply func, not calling it. i.e., why not parentheses
	default:
		return func(f1, f2 float64) float64 {
			return 0
		}
	}
}

// func as args
func double(f1, f2 float64, mathExpr func(float64, float64) float64) float64 {
	return 2 * mathExpr(f1, f2)
}
