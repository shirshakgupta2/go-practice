// BFSExample demonstrates a basic implementation of the Breadth-First Search (BFS) algorithm.
// It can be used to traverse or search through graph structures, such as trees or networks.	
/*
Use Case: Finding the shortest path in an unweighted graph.
*/


package main

import (
	"container/list"
	"fmt"
)

type Graph struct {
	adj     map[int][]int
	directed bool
}

func NewGraph(directed bool) *Graph {
	return &Graph{
		adj:      make(map[int][]int),
		directed: directed,
	}
}

func (g *Graph) AddEdge(from, to int) {
	g.adj[from] = append(g.adj[from], to)
	if !g.directed {
		g.adj[to] = append(g.adj[to], from)
	}
}

func (g *Graph) BFS(start int) {
	visited := make(map[int]bool)
	queue := list.New()

	queue.PushBack(start)
	visited[start] = true

	for queue.Len() > 0 {
		curr := queue.Remove(queue.Front()).(int)
			fmt.Printf("%d ", curr)

		for _, neighbor := range g.adj[curr] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue.PushBack(neighbor)
			}
		}
	}
	fmt.Println()
}

// DFSExample demonstrates a basic implementation of the Depth-First Search (DFS) algorithm.
// It can be used to traverse or search through graph structures, such as trees or networks.
// Use Case: Finding all connected components in a graph.
// It can also be used to explore all paths in a maze or puzzle.

func (g *Graph) DFS(start int) {
	visited := make(map[int]bool)
	g.dfsHelper(start, visited)
	fmt.Println()
}

func (g *Graph) dfsHelper(node int, visited map[int]bool) {
	visited[node] = true
	fmt.Printf("%d ", node)

	for _, neighbor := range g.adj[node] {
		if !visited[neighbor] {
			g.dfsHelper(neighbor, visited)
		}
	}
}
