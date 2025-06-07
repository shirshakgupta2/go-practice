package main

import (
	"container/list"
	"fmt"
)


func (graph *Graph) validPath(n int, edges [][]int, source int, destination int) bool {
	for i := range edges {
		graph.AddEdge(edges[i][0], edges[i][1], true)
	}
	graph.Print()
	visitedMap := make(map[int]bool, n)
	queue := list.New()
	queue.PushBack(source)
	visitedMap[source] = true

	for queue.Len() > 0 {
		ele := queue.Front()
		curr := queue.Remove(ele).(int)
		for _, adjEdge := range graph.adj[curr] {
			if !visitedMap[adjEdge] {
				if adjEdge == destination {
					return true
				}
				visitedMap[adjEdge] = true
				queue.PushBack(adjEdge)
			}
		}
		if visitedMap[destination] {
			return visitedMap[destination]
		}

	}
	return visitedMap[destination]
	// return false
}

func (g *Graph) Print() {
	for node, neighbors := range g.adj {
		fmt.Printf("%d -> %v\n", node, neighbors)
	}
}

// func validPathBFS() {

// }
