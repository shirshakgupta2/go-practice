package main

import "fmt"

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	key       int
	adjacents []*Vertex
}

func (g *Graph) AddVertex(key int) *Vertex {
	vertex := &Vertex{key: key}
	// Check if the vertex already exists
	if containsVertex(g.vertices, key) {
		return nil
	}
	g.vertices = append(g.vertices, vertex)
	return vertex
}

func (g *Graph) GetVertex(key int) *Vertex {
	for _, v := range g.vertices {
		if v.key == key {
			return v
		}
	}
	return nil
}

func containsVertex(vertices []*Vertex, key int) bool {
	for _, v := range vertices {
		if v.key == key {
			return true
		}
	}
	return false
}

func containsVertexInAdjacencyList(adjacents []*Vertex, key int) bool {
	for _, v := range adjacents {
		if v.key == key {
			return true
		}
	}
	return false
}

//Directed graph edge addition
// This function adds an edge from one vertex to another
func (g *Graph) AddEdge(from, to int) bool {
	fromVertex := g.GetVertex(from)
	toVertex := g.GetVertex(to)

	if fromVertex == nil || toVertex == nil {
		fmt.Printf("One or both vertices not found: %d -> %d\n", from, to)
		return false
	}
	// Check if the edge already exists
	if containsVertexInAdjacencyList(fromVertex.adjacents, to) {
		fmt.Printf("Edge already exists: %d -> %d\n", from, to)
		return false
	}

	fromVertex.adjacents = append(fromVertex.adjacents, toVertex)
	return true
}

func AdvanceGraphAdjacencyList() {
	g := &Graph{}
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)
	g.AddVertex(3)

	//basic adjacency setup
	// v1.adjacents = append(v1.adjacents, v2, v3)
	// v2.adjacents = append(v2.adjacents, v3)
	// fmt.Println("Graph created with vertices and adjacencies:", g.vertices)

	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(1, 3)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.AddEdge(4, 1)
	g.AddEdge(4, 1)
	g.AddEdge(2, 4)
	fmt.Println("Graph created with vertices and edges:")

	for _, v := range g.vertices {
		fmt.Println("Vertex:", v.key)
		for _, adj := range v.adjacents {
			fmt.Println("  Adjacent to:", adj.key)
		}
	}
}


