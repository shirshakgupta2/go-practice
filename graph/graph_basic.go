package main

import "fmt"

type SimpleGraph struct {
	adj map[int][]int
}

func NewGraph() *SimpleGraph {
	return &SimpleGraph{adj: make(map[int][]int)}
}

func (g *SimpleGraph) AddEdge(u, v int, directed bool) {
	g.adj[u] = append(g.adj[u], v)
	if !directed {
		g.adj[v] = append(g.adj[v], u)
	}
}

func (g *SimpleGraph) Print() {
	for node, neighbors := range g.adj {
		fmt.Printf("%d -> %v\n", node, neighbors)
	}
}


func BasicGraphAdjacencyList() {
	fmt.Println("Basic Graph Example:")
	g := NewGraph()
	g.AddEdge(1, 2, false)
	g.AddEdge(1, 3, false)
	g.AddEdge(2, 3, false)
	g.AddEdge(3, 4, true)
	g.AddEdge(4, 1, false)
	g.Print()
}