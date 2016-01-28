package main

import (
	"fmt"
	"github.com/iyoo14/goalgo/lib"
)

const (
	WHITE = iota
	GRAY
	BLACK
	N     = 10
	INFTY = 1 << 21
)

var v, n, tt int
var M [N][N]int
var d [N]int

func bfs(s int) {
	q := lib.NewQueue(N)
	q.Enqueue(s)
	for i := 0; i < n; i++ {
		d[i] = INFTY
	}
	d[s] = 0
	for q.IsEmpty() {
		u, _ := q.Dequeue()
		iu := u.(int)
		for v = 0; v < n; v++ {
			if M[iu][v] == 0 {
				continue
			}
			if d[v] != INFTY {
				continue
			}
			d[v] = d[iu] + 1
			q.Enqueue(v)
		}
	}
	for i := 0; i < n; i++ {
		if d[i] == INFTY {
			d[i] = -1
		}
		fmt.Printf("%d %d\n", i+1, d[i])
	}
}

func main() {

	var u, k, v int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			M[i][j] = 0
		}
	}
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &u)
		fmt.Scanf("%d", &k)
		u--
		for j := 0; j < k; j++ {
			fmt.Scanf("%d", &v)
			v--
			M[u][v] = 1
		}
	}
	bfs(0)
}
