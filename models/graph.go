// Package models/graph defines the graph data structure for use in this module.
package models

// Node is the building block of graph type in goGraph is built on the node type.
//
// in the node type we store data about the node and pointers to the nodes
// neighbors. Every Node knows about its neighbors
type Node struct {
	Idx          int               // node index
	Neighborhood map[*Node]float64 // neighbors of the node and their edge weights as map
	Weight       float64           // Node Weight
	Attributes   map[string]string // memory available for other attributes.
}

func (n Node) Degree() int {
	return len(n.Neighborhood)
}

func (n Node) IsIsolated() bool {
	switch n.Degree() {
	case 0:
		return true
	default:
		return false
	}
}

func (n Node) IsLeaf() bool {
	switch n.Degree() {
	case 1:
		return true
	default:
		return false
	}
}

// Graph type
type Graph struct {
	Nodes        map[int]Node
	Directed     bool
	EdgeWeighted bool
	NodeWeighted bool
}
