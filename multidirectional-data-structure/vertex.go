package directed_graph

import "fmt"

// Type of a vertex in a directed graph
type DirectedVertex struct {
	key          int
	OutNeighbors []*DirectedVertex
	InNeighbors  []*DirectedVertex
}

// AddVertex adds a node. This increases the space used by the adjacency matrix
// at a quadratic rythm, manner, so beware of memory usage.
func (dg *DirectionalGraph) AddVertex(k int) {
	// checks to see if the key is already in use
	for _, node := range dg.vertices {
		if k == node.key {
			err := fmt.Errorf("Vertex %v not added because it uses an existing key", k)
			fmt.Println(err.Error())
		}
	}
	dg.vertices = append(dg.vertices, &DirectedVertex{key: k})
}

// PruneVertex removes a node regardless of wether it has any edges or not.
// If it has, the vertices are also removed.
func (dg *DirectionalGraph) PruneVertex(k int) {

}

// RemoveVertex removes a node, but only if the node has no edges. To
// remove regardless of the edges, use PruneVertex().
func (dg *DirectionalGraph) RemoveVertex(k int) {

}

// GetVertex
func (dg *DirectionalGraph) GetVertex(k int) *DirectedVertex {
	for _, node := range dg.vertices {
		if k == node.key {
			return node
		}
	}
	return nil
}
