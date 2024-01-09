package basic_graph

import "fmt"

// Graph represents an adjacency list graph
type Graph struct {
	vertices []*Vertex
}

// Print will print the adjacency list for all vertices of the graph
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v : ", v.key)
		for _, v := range v.adjacent {
			fmt.Printf(" %v ", v.key)
		}
	}
	fmt.Println()
}
