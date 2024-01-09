package basic_graph

import "fmt"

// Vertex represents a graph vertex
type Vertex struct {
	key      int
	adjacent []*Vertex
}

// AddVertex adds a vertex to the graph with a provided key. If the key is already
// taken, the operation fails, gives off a warning but does not stop the program.
func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("Vertex %v not added because it uses an existing key", k)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}

// GetVertex returns checks to see if there is a vertex with the parameter key (k).
// If it succeeds in finding, it returns the vertex. If it doesn't, it returns nil.
func (g *Graph) GetVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

// contains tells wether or not a vertex (s) has a certain key (k)
func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}
