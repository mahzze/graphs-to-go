package directed_graph

import "fmt"

// The directional graph is  weighted and represented by a an adjacency matrix.
type DirectionalGraph struct {
	nodes []*DirectedNode
	edges [][]int
}

func (dg *DirectionalGraph) Print() {
	for _, v := range dg.nodes {
		fmt.Printf("= node key: %v -- inNeighbors : %v\n= outNeighbors: %v\n", v.key, v.InNeighbors, v.OutNeighbors)
	}
}
