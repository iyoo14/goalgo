package main

import (
	"fmt"
	"github.com/iyoo14/goalgo/lib"
)

const (
	MAX = 10
	NIL = -1
)

var n int
var G [MAX]lib.Vector
var color [MAX]int

func dfs(r int, c int) {
	s := lib.NewStack(MAX)
	s.Push(r)
	color[r] = c
	for !s.IsEmpty() {
		u := s.Top()
		s.Pop()
		gu := G[u.(int)]
		for i := 0; i < gu.Size(); i++ {
			v := gu.V[i]
			if color[v] == NIL {
				color[v] = c
				s.Push(v)
			}
		}
	}
}

func assignColor() {
	id := 1
	for i := 0; i < n; i++ {
		color[i] = NIL
	}
	for u := 0; u < n; u++ {
		if color[u] == NIL {
			dfs(u, id)
			id++
		}
	}
}

func main() {
	var s, t, m, q int

	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &m)

	for i := 0; i < n; i++ {
		G[i] = lib.NewVector(MAX)
	}
	for i := 0; i < m; i++ {
		fmt.Scanf("%d", &s)
		fmt.Scanf("%d", &t)
		vs := G[s]
		vs.PushBack(t)
		vt := G[t]
		vt.PushBack(s)
	}

	fmt.Println(G)

	assignColor()

	fmt.Scanf("%d", &q)

	fmt.Println(color)
	for i := 0; i < q; i++ {
		fmt.Scanf("%d", &s)
		fmt.Scanf("%d", &t)
		if color[s] == color[t] {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}
