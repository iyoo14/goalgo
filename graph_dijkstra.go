package main

import (
	"fmt"
)

const (
	WHITE = iota
	GRAY
	BLACK
	MAX   = 10
	INFTY = 1 << 21
)

var M [MAX][MAX]int
var n int

func dijkstra() {
	var minv int
	var d, color [MAX]int

	for i := 0; i < n; i++ {
		d[i] = INFTY
		color[i] = WHITE
	}

	d[0] = 0
	color[0] = GRAY
	for {
		minv = INFTY
		u := -1
		for i := 0; i < n; i++ {
			if minv > d[i] && color[i] != BLACK {
				u = i
				minv = d[i]
			}
		}
		if u == -1 {
			break
		}
		color[u] = BLACK
		for v := 0; v < n; v++ {
			if color[v] != BLACK && M[u][v] != INFTY {
				if d[v] > d[u]+M[u][v] {
					d[v] = d[u] + M[u][v]
					color[v] = GRAY
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		if d[i] == INFTY {
			d[i] = -1
		}
		fmt.Printf("%d %d\n", i, d[i])
	}
}

func main() {

	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		for j := 0; j < 0; j++ {
			M[i][j] = INFTY
		}
	}

	var k, c, u, v int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &u)
		fmt.Scanf("%d", &k)
		for j := 0; j < k; j++ {
			fmt.Scanf("%d", &v)
			fmt.Scanf("%d", &c)
			M[u][v] = c
		}
	}
	dijkstra()
}
