package main

import "fmt"

type Graph struct {
	adj map[int][]int
}

func NewGraph() *Graph {
	return &Graph{
		adj: make(map[int][]int),
	}

}

func (graph *Graph) AddEdge(from, to int, bidirectional bool) {
	// graph.adj[from][len(graph.adj[from])] = to
	graph.adj[from] = append(graph.adj[from], to)
	if bidirectional {
		graph.adj[to] = append(graph.adj[to], from)
	}
}

func main() {
	//!FIND MINIMUM STEPS TO REACH TARGET POSITION FOR KNIGHT
	knight := Knight{Position: Position{X: 1, Y: 3}}
	target := Position{X: 5, Y: 0}
	stepsWithBFS := knight.minimumStepsToReachBFS(target)
	fmt.Printf("Minimum steps to reach target (%d, %d): %d\n", target.X, target.Y, stepsWithBFS)

	//!FIND IF PATH EXISTS
	graph := NewGraph()
	noOfEdges:=10
	edges := [][]int{
		{4, 3}, {1, 4}, {4, 8}, {1, 7}, {6, 4},
		{4, 2}, {7, 4}, {4, 0}, {0, 9}, {5, 4},
	}
	source:=5
	destination:=9
	fmt.Println("Source to destination path DoesExists:", graph.validPath(noOfEdges, edges, source, destination))

	//!FIND JUDGE
	n := 3
	trusts := [][]int{
		{1, 3},
		{2, 3},
		{3, 1},
	}
	judgegraph := NewGraph()
	fmt.Println("Judge is:", judgegraph.findJudge(n, trusts))


	//! NO OF ISLAND
	fmt.Println("NoOfIsland:",NoOfIsland())


}
