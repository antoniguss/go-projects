package utils

import (
	"fmt"
	"math"
)

// PrintMatrix prints a  2D array in a formatted manner
func PrintMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, value := range row {
			if value == math.MaxInt {
				fmt.Print("∞ ") // Print infinity for max int
			} else {
				fmt.Printf("%d ", value)
			}
		}
		fmt.Println() // Newline after each row
	}
}

func Create2DArray(a, b int) ([][]int, error) {
	if a <= 0 || b <= 0 {
		return nil, fmt.Errorf("Invalid array dimensions: %dx%d", a, b)
	}

	matrix := make([][]int, a)
	for i := range a {
		matrix[i] = make([]int, b)
	}

	return matrix, nil
}
