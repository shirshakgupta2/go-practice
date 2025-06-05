package main

import "fmt"

func main() {
	g := NewGraph(true) // change to `true` for directed graph

	edges := [][2]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {3, 6}}

	for _, e := range edges {
		g.AddEdge(e[0], e[1])
	}

	fmt.Println("Graph Adjacency List:")
	for node, neighbors := range g.adj {
		fmt.Printf("%d -> %v\n", node, neighbors)
	}

	fmt.Println("BFS from node 0:")
	g.BFS(0)

	fmt.Println("DFS from node 0:")
	g.DFS(0)

}
