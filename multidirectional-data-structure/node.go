package directed_graph

import "fmt"

// Type of a vertex in a directed graph
type DirectedNode struct {
	key          int
	OutNeighbors []*DirectedNode
	InNeighbors  []*DirectedNode
}

// AddNode adds a node. This increases the space used by the adjacency matrix
// at a quadratic rythm, manner, so beware of memory usage.
func (dg *DirectionalGraph) AddNode(k int) {
	// checks to see if the key is already in use
	for _, node := range dg.nodes {
		if k == node.key {
			err := fmt.Errorf("Vertex %v not added because it uses an existing key", k)
			fmt.Println(err.Error())
		}
	}
	dg.nodes = append(dg.nodes, &DirectedNode{key: k})
}

// PruneNode removes a node regardless of wether it has any edges or not.
// If it has, the nodes are also removed.
func (dg *DirectionalGraph) PruneNode(k int) {

}

// RemoveNode removes a node, but only if the node has no edges. To
// remove regardless of the edges, use PruneNode().
func (dg *DirectionalGraph) RemoveNode(k int) {

}

// GetNode
func (dg *DirectionalGraph) GetNode(k int) *DirectedNode {
	for _, node := range dg.nodes {
		if k == node.key {
			return node
		}
	}
	return nil
}
