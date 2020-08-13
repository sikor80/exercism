// Package diffsquares is a solution to exercism.io exercise titled Difference Of Squares.
package diffsquares

// SquareOfSum returns a square of sum of first N natural numbers.
func SquareOfSum(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum * sum
}

// SumOfSquares returns a sum of squares of first N natural numbers.
func SumOfSquares(n int) int {
	var squareSum int
	for i := 0; i <= n; i++ {
		squareSum += i * i
	}

	return squareSum
}

// Difference returns a difference between SquareOfSum and SumOfSquares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
