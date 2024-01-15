package directed_graph

import "fmt"

// The directional graph is  weighted and represented by a an adjacency matrix.
type DirectionalGraph struct {
	vertices []*DirectedVertex
	edges    [][]int
}

func (dg *DirectionalGraph) PrintVertices() {
	for _, v := range dg.vertices {
		fmt.Printf("= vertex key: %v -- inNeighbors : %v\n= outNeighbors: %v\n", v.key, v.InNeighbors, v.OutNeighbors)
	}
}
