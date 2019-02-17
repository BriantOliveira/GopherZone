package diffsquares

// SquareOfSum calculates the square of the sum of all numbers smaller or equal to the input
func SquareOfSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

// SumOfSquares calculates the sum of squares of all numbers smaller or equal to the input
func SumOfSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

// Difference calculates the difference between the square of sums and the sum of squares of all numbers
// smaller or equal to the input
// More efficient than calling the two functions above since it only requires one for loop
func Difference(n int) int {
	sum1 := 0
	sum2 := 0
	for i := 1; i <= n; i++ {
		sum1 += i * i
		sum2 += i
	}
	result := sum2*sum2 - sum1
	return result
}
