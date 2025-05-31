package patterns

import (
	"fmt"
)

// func PascalsTriangle(rows int) {
// 	for i := 0; i < rows; i++ {
// 		number := 1

// 		// Print leading spaces for formatting
// 		fmt.Printf("%*s", (rows-i)*2, "")

// 		// Print numbers in the row
// 		for j := 0; j <= i; j++ {
// 			fmt.Printf("%4d", number)
// 			number = number * (i - j) / (j + 1)
// 		}
// 		fmt.Println()
// 	}
// }

// Function to compute n choose k
func binomialCoeff(n, k int) int {
	res := 1
	if k > n-k {
		k = n - k
	}
	for i := 0; i < k; i++ {
		res = res * (n - i) / (i + 1)
	}
	return res
}

func PascalsTriangle(rows int) {
	// Determine the maximum number to compute its width
	maxNum := binomialCoeff(rows-1, (rows-1)/2)
	width := len(fmt.Sprintf("%d", maxNum)) + 1 // Add space between numbers

	for i := 0; i < rows; i++ {
		number := 1

		// Print leading spaces
		fmt.Printf("%*s", (rows-i)*width/2, "")

		for j := 0; j <= i; j++ {
			fmt.Printf("%*d", width, number)
			number = number * (i - j) / (j + 1)
		}
		fmt.Println()
	}
}
