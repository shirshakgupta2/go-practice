package main

import "fmt"

type WeightedGraph struct {
	weighted_vertices []*WeightedVertex
}

type WeightedVertex struct {
	key       int
	adjacents []*WeightedEdge
}

type WeightedEdge struct {
	to     *WeightedVertex
	weight int
}

func (g *WeightedGraph) AddVertex(key int) *WeightedVertex {
	vertex := &WeightedVertex{key: key}
	// Check if the vertex already exists
	if containsWeightedVertex(g.weighted_vertices, key) {
		return nil
	}
	g.weighted_vertices = append(g.weighted_vertices, vertex)
	return vertex
}

func (g *WeightedGraph) AddEdge(from, to int, weight int) bool {
	fromVertex := g.GetVertex(from)
	toVertex := g.GetVertex(to)

	if fromVertex == nil || toVertex == nil {
		fmt.Printf("One or both vertices not found: %d -> %d\n", from, to)
		return false
	}

	// Check if the edge already exists
	if containsEdge(fromVertex.adjacents, to) {
		fmt.Printf("Edge already exists: %d -> %d\n", from, to)
		return false
	}

	edge := &WeightedEdge{to: toVertex, weight: weight}
	fromVertex.adjacents = append(fromVertex.adjacents, edge)
	return true
}

func containsWeightedVertex(vertices []*WeightedVertex, key int) bool {
	for _, v := range vertices {
		if v.key == key {
			return true
		}
	}
	return false
}


func containsEdge(adjacents []*WeightedEdge, to int) bool {
	for _, edge := range adjacents {
		if edge.to.key == to {
			return true
		}
	}
	return false
}

func (g *WeightedGraph) GetVertex(key int) *WeightedVertex {
	for _, v := range g.weighted_vertices {
		if v.key == key {
			return v
		}
	}
	return nil
}

func WeightedGraphAdjacencyList() {

	fmt.Println("Weighted Graph Example:")
	g := &WeightedGraph{}
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)

	g.AddEdge(1, 2, 5)
	g.AddEdge(1, 3, 10)
	g.AddEdge(2, 3, 2)
	g.AddEdge(3, 4, 1)
	g.AddEdge(4, 1, 3)

	for _, vertex := range g.weighted_vertices {
		fmt.Printf("Vertex %d: ", vertex.key)
		for _, edge := range vertex.adjacents {
			fmt.Printf(" -> %d (weight: %d)", edge.to.key, edge.weight)
		}
		fmt.Println()
	}
}
