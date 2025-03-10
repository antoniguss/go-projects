package graphs

import (
	"fmt"
	"math"

	"algorithms/utils"
)

const Infinity = math.MaxInt64

func RunBellmanFordExample() {
	gm := NewGraphMatrix(5)

	gm.AddEdge(0, 1, 6)  // Edge from node 0 to 1 with weight 6
	gm.AddEdge(0, 2, 7)  // Edge from node 0 to 2 with weight 7
	gm.AddEdge(1, 2, 8)  // Edge from node 1 to 2 with weight 8
	gm.AddEdge(1, 3, 5)  // Edge from node 1 to 3 with weight 5
	gm.AddEdge(1, 4, -4) // Edge from node 1 to 4 with weight -4
	gm.AddEdge(2, 3, -3) // Edge from node 2 to 3 with weight -3
	gm.AddEdge(3, 1, 2)  // Edge from node 3 to 1 with weight 2
	gm.AddEdge(4, 0, 2)  // Edge from node 4 to 0 with weight 2
	gm.AddEdge(4, 3, 7)  // Edge from node 4 to 3 with weight 7

	utils.PrintMatrix(gm.matrix)

	fmt.Println()

	L := RunBellmanFord(*gm)

	utils.PrintMatrix(L)
}

// Dynamic programming implementation
func RunBellmanFord(gm GraphMatrix) [][]int {
	n := len(gm.matrix)

	L, _ := utils.Create2DArray(n, n)

	for i := 0; i < n; i++ {
		L[i][0] = Infinity
	}
	L[0][0] = 0

	for m := 1; m < n; m++ {
		for k := 0; k < n; k++ {

			minFound := Infinity
			for j := 0; j < n; j++ {
				if L[j][m-1] != Infinity {
					weight := gm.Weight(j, k)

					if weight != Infinity {
						minFound = min(minFound, L[j][m-1]+weight)
					}

				}
			}

			L[k][m] = minFound
		}
	}

	return L
}
