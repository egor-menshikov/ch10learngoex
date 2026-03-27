// Package ch10learngoex is a chapter 10 exercise package from "Learning Go" 2nd ed by Jon Bodner
package ch10learngoex

import "golang.org/x/exp/constraints"

// Number is a constraint that permits any integer or floating-point type.
type Number interface {
	constraints.Float | constraints.Integer
}

// Add adds two integers together and returns the result.
// For more information on addition, see https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}
