package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("fuck")
	lengths := []int{3, 9, 4, 1, 5, 4, 2, 4, 2, 2, 10, 7}

	maxLen := 18

	results := tidy_typesetting(lengths, maxLen)

	for i, v := range results {
		fmt.Printf("i: %d, T[i] = %d\n", i, v)
	}
}

func tidy_typesetting(l []int, L int) (T []int) {
	n := len(l)
	l = append([]int{0}, l...) // For the index to start at 1
	T = make([]int, n+1)

	T[0] = 0

	// Precompute `w` and `c` values
	w := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		w[i] = make([]int, n+1)
	}

	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			if i == j {
				w[i][j] = l[i]
				continue
			}

			w[i][j] = w[i][j-1] + l[j] + 1
		}
	}

	c := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		c[i] = make([]int, n+1)
	}

	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {

			width := w[i][j]
			if width > L {
				continue
			}

			cost := (L - width) * (L - width) * (L - width)

			c[i][j] = cost
		}
	}

	printGrid(w)
	printGrid(c)

	for j := 1; j <= n; j++ {
		minCost := math.MaxInt
		minFor := 0

		for i := j; i >= 1; i-- {

			width := w[i][j]
			if width > L {
				break
			}

			if c[i][j]+T[i-1] < minCost {
				minFor = i
			}
			minCost = min(minCost, c[i][j]+T[i-1])
		}

		fmt.Printf("j = %d, minFor := %d\n", j, minFor)
		T[j] = minCost

	}

	return T
}

func printGrid(matrix [][]int) {
	if len(matrix) == 0 {
		fmt.Println("The matrix is empty.")
		return
	}

	for _, row := range matrix {
		for _, value := range row {
			fmt.Printf("%4d ", value) // Adjust spacing as needed
		}
		fmt.Println()
	}
}
