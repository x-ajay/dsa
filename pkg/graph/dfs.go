package graph

import "fmt"

func dfs(graph [][]int, visited []bool, source int) {
	visited[source] = true
	fmt.Printf("%d->", source)
	for i := 0; i < len(graph[0]); i++ {
		if graph[source][i] == 1 && !visited[i] {
			visited[i] = true
			dfs(graph, visited, i)
		}
	}
}

func DFS(graph [][]int, source int) {
	visited := make([]bool, len(graph[0]))
	dfs(graph, visited, source)
	fmt.Printf("done\n")
}
