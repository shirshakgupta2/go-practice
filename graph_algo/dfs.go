package main

import "fmt"

func (g *Graph) testDFS(source int, vertex int) {
	
	visitedMap := make(map[int]bool)
	// for i := range vertex {
	// 	if !visitedMap[i] {
	// 		g.checkDFS(i, visitedMap)
	// 	}
	// }
	g.checkDFS(source, visitedMap)
}

func (g *Graph) checkDFS(i int, visitedMap map[int]bool) {
	visitedMap[i] = true
	fmt.Printf("%d ", i)
	for _, adj := range g.adj[i] {
		if !visitedMap[adj] {
			g.checkDFS(adj, visitedMap)
		}
	}
}
