package graphs

import (
	"math"
)

type GraphMatrix struct {
	matrix [][]int
}

func NewGraphMatrix(n int) *GraphMatrix {
	matrix := make([][]int, n)

	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	for u := 0; u < n; u++ {
		for v := 0; v < n; v++ {
			if u == v {
				matrix[u][v] = 0
			} else {
				matrix[u][v] = math.MaxInt
			}
		}
	}

	return &GraphMatrix{matrix: matrix}
}

func (gm *GraphMatrix) AddEdge(u, v int, weight int) bool {
	if u >= len(gm.matrix) || v >= len(gm.matrix) || u == v {
		return false
	}

	gm.matrix[u][v] = weight

	return true
}

func (gm *GraphMatrix) Weight(u, v int) int {
	if u >= len(gm.matrix) || v >= len(gm.matrix) {
		return math.MaxInt
	}

	if u == v {
		return 0
	}

	return gm.matrix[u][v]
}
