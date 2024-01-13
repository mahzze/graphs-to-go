package directed_graph

// Type of a vertex in a directed graph
type DirectedNode struct {
	key          int
	OutNeighbors []*DirectedNode
	InNeighbors  []*DirectedNode
}

// AddNode adds a node. It works by copying the adjacency matrix into a new matrix
// with the new node and returning the new matrix, so beware of memory usage.

// PruneNode removes a node regardless of wether it has any edges or not.

// RemoveNode removes a node, but only if the node has no edges. To
// remove regardless of the edges, use PruneNode().

// GetNode
func (g *DirectionalGraph) GetNode(k int) *DirectedNode {
	for x := range g.nodes {
		for y := range g.nodes[x] {
			if k == g.nodes[x][y].key {
				return g.nodes[x][y]
			}
		}
	}
	return nil
}
