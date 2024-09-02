package graph

import (
	"fmt"

	"github.com/x-ajay/dsa/pkg/queue"
)

/*	0-> 1,2,4
	1-> 0,4
	2-> 0,3
	3-> 2,4,5
	4-> 0,1,3,5
	5-> 3,4

	mat := [][]int{
		{0, 1, 1, 0, 1, 0},
		{1, 0, 0, 0, 1, 0},
		{1, 0, 0, 1, 0, 0},
		{0, 0, 1, 0, 1, 1},
		{1, 1, 0, 1, 0, 1},
		{0, 0, 0, 1, 1, 0},
	}
*/

func BFS(graph [][]int, source int) {
	visited := make([]bool, len(graph[0]))

	visited[source] = true
	q := queue.New()
	q.Enqueue(source)

	for !q.IsEmpty() {
		front := q.Front()
		q.Dequeue()

		fmt.Printf("%d->", front)

		for i := 0; i < len(graph[front]); i++ {
			if graph[front][i] == 1 && !visited[i] {
				q.Enqueue(i)
				visited[i] = true
			}
		}
	}
	fmt.Printf("done\n")
}
