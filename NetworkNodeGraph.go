package main

import "fmt"

type NetworkNodeGraph struct {
	vertices []NetworkNode
}

func (graph NetworkNodeGraph) printEdges() {
	for _, node := range graph.vertices {
		for _, dst := range node.neighbors {
			fmt.Println(node.name, dst)
		}
	}
}

func (graph *NetworkNodeGraph) addEdge(srcNode string, dstNode string) {
	nodeRef := graph.getNode(srcNode)
	if graph.contains(dstNode) {
		if !nodeRef.hasNeighbor(dstNode) {
			nodeRef.neighbors = append(nodeRef.neighbors, dstNode)
		}
	} else {
		graph.addVertex(dstNode)
		nodeRef.neighbors = append(nodeRef.neighbors, dstNode)
	}
}

func (graph *NetworkNodeGraph) getNode(nodeName string) *NetworkNode {
	for i := 0; i < len(graph.vertices); i++ {
		if graph.vertices[i].name == nodeName {
			return &graph.vertices[i]
		}
	}
	return &NetworkNode{}
}

func (graph *NetworkNodeGraph) addVertex(nodeName string) {
	node := new(NetworkNode)
	node.name = nodeName
	graph.vertices = append(graph.vertices, *node)
}

func (graph *NetworkNodeGraph) contains(nodeName string) bool {
	for _, node := range graph.vertices {
		if node.name == nodeName {
			return true
		}
	}
	return false
}
