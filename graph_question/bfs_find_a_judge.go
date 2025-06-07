package main

func (graph *Graph) findJudge(n int, trust [][]int) int {
	indegreeMap:= make(map[int]int)
	outdegreeMap:= make(map[int]int)
	for i := range trust {
		
		graph.AddEdge(trust[i][0], trust[i][1], false)
		indegreeMap[trust[i][1]]++ 
		outdegreeMap[trust[i][0]]++
	}
	graph.Print()
	for i:=1;i<=n;i++{
		if indegreeMap[i]==n-1 &&outdegreeMap[i]==0{
			return i
		}
	}
	return -1
 

}
