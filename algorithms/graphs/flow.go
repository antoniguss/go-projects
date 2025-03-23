package graphs

import "fmt"

func RunFlowExample() {
	fmt.Println("Calculating max flow")
	adjacencyList := map[int][]int{
		1: {2, 4},
		2: {3, 5, 1},
		3: {6, 2},
		4: {1, 5, 7},
		5: {2, 6, 8, 4},
		6: {3, 0, 9, 5},
		7: {4, 8},
		8: {5, 9, 7},
		9: {6, 0, 8},
	}

	g := NewGraph(WithAdjacencyList(adjacencyList))
}

// Graph representation adapted from https://gist.github.com/snassr/e79f4953eeb8d813be82eda00adeef57
type Graph struct {
	Vertices map[int]*Vertex
}

type Vertex struct {
	Val   int
	Edges map[int]*Edge
}

type Edge struct {
	Weight int
	Vertex *Vertex
}

func NewGraph(opts ...GraphOption) *Graph {
	g := &Graph{Vertices: map[int]*Vertex{}}

	for _, opt := range opts {
		opt(g)
	}

	return g
}

type GraphOption func(this *Graph)

func WithAdjacencyList(list map[int][]int) GraphOption {
	return func(this *Graph) {
		for vertex, edges := range list {
			if _, has := this.Vertices[vertex]; !has {
				this.AddVertex(vertex, vertex)
			}

			for _, edge := range edges {
				if _, has := this.Vertices[edge]; !has {
					this.AddVertex(edge, edge)
				}

				this.AddEdge(vertex, edge, 0)
			}
		}
	}
}

func (this *Graph) AddVertex(key, val int) {
	this.Vertices[key] = &Vertex{Val: val, Edges: map[int]*Edge{}}
}

func (this *Graph) AddEdge(srcKey, destKey int, weight int) {
	if _, has := this.Vertices[srcKey]; !has {
		return
	}
	if _, has := this.Vertices[destKey]; !has {
		return
	}

	this.Vertices[srcKey].Edges[destKey] = &Edge{
		Weight: weight,
		Vertex: this.Vertices[destKey],
	}
}

func (this *Graph) RemoveEdge(srcKey, destKey int, weight int) {
	if _, has := this.Vertices[srcKey]; !has {
		return
	}
	if _, has := this.Vertices[destKey]; !has {
		return
	}

	delete(this.Vertices[srcKey].Edges, destKey)
}

func (this *Graph) Neighbours(srcKey int) []int {
	result := []int{}

	if _, has := this.Vertices[srcKey]; !has {
		return result
	}

	for _, edge := range this.Vertices[srcKey].Edges {
		result = append(result, edge.Vertex.Val)
	}

	return result
}

func (this *Graph) GetEdge(srcKey, destKey int) (edge *Edge, has bool) {
	if _, has := this.Vertices[srcKey]; !has {
		return &Edge{}, false
	}

	edge, has = this.Vertices[srcKey].Edges[destKey]
	return edge, has
}
