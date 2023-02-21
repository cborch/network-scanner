package main

type NetworkNode struct {
	name      string
	neighbors []string
}

// Just using simple (linear) search for now
func (node NetworkNode) hasNeighbor(nodeName string) bool {
	for _, neighbor := range node.neighbors {
		if neighbor == nodeName {
			return true
		}
	}
	return false
}
