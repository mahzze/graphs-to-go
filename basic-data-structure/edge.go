package basic_graph

import "fmt"

// There is not an actual edge structure in adjacency lists and matrices!
// But there are edge-related methods, so they will be contained in their
// own file.

// Adds an edge between the two vertices with the keys A and B. If there
// is no corresponding vertex for the keys, returns an error.
func (g *Graph) AddEdge(keyA, keyB int) {
	// get vertex
	vertexA := g.GetVertex(keyA)
	vertexB := g.GetVertex(keyB)

	// check for errors
	if vertexA == nil || vertexB == nil {
		err := fmt.Errorf("Invalid edge (%v <--> %v)", vertexA, vertexB)
		fmt.Println(err.Error())
	} else if contains(vertexA.adjacent, keyB) || contains(vertexB.adjacent, keyA) {
		err := fmt.Errorf("Already existing edge (%v <--> %v)", vertexA, vertexB)
		fmt.Println(err.Error())
	} else {
		// add edges to both, for this graph is not a directional one.
		vertexA.adjacent = append(vertexA.adjacent, vertexB)
		vertexB.adjacent = append(vertexB.adjacent, vertexA)
	}
}
