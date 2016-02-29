package main

import (
	"fmt"
	"sort"
)

type Node struct {
	location int
	p        int
	l        int
	r        int
}

type Point struct {
	id int
	x  int
	y  int
}

func NewPoint(id int, x int, y int) Point {
	p := Point{id, x, y}
	return p
}

type Ps []Point

func (p Ps) Len() int {
	return len(p)
}

func (p Ps) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *Point) disp() {
	fmt.Printf("%d\n", p.id)
}

type ByX struct {
	Ps
}

type ByY struct {
	Ps
}

type ById struct {
	Ps
}

func (b ByX) Less(i, j int) bool {
	return b.Ps[i].x < b.Ps[j].x
}

func (b ByY) Less(i, j int) bool {
	return b.Ps[i].y < b.Ps[j].y
}

func (b ById) Less(i, j int) bool {
	return b.Ps[i].id < b.Ps[j].id
}

const (
	NIL = -1
	MAX = 10
)

var N int
var P Ps
var T [MAX]Node
var np int

func makeKDTree(l int, r int, depth int) int {
	if !(l < r) {
		return NIL
	}
	mid := (l + r) / 2
	t := np
	np++
	if (depth % 2) == 0 {
		sort.Sort(ByX{P[l:r]})
	} else {
		sort.Sort(ByY{P[l:r]})
	}
	T[t].location = mid
	T[t].l = makeKDTree(l, mid, depth+1)
	T[t].r = makeKDTree(mid+1, r, depth+1)

	return t
}

func find(v int, sx int, tx int, sy int, ty int, depth int, ans Ps) {
	x := P[T[v].location].x
	y := P[T[v].location].y

	if sx <= x && x <= tx && sy <= y && y <= ty {
		ans = append(ans, P[T[v].location])
	}

	if depth%2 == 0 {
		if T[v].l != NIL {
			if sx <= x {
				find(T[v].l, sx, tx, sy, ty, depth+1, ans)
			}
		}
		if T[v].r != NIL {
			if x <= tx {
				find(T[v].r, sx, tx, sy, ty, depth+1, ans)
			}
		}
	} else {
		if T[v].l != NIL {
			if sy <= y {
				find(T[v].l, sx, ty, sy, ty, depth+1, ans)
			}
		}
		if T[v].r != NIL {
			if y <= ty {
				find(T[v].r, sx, ty, sy, ty, depth+1, ans)
			}
		}
	}
}

func main() {
	var x, y int

	fmt.Scanf("%d", &N)
	P = make(Ps, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d %d", &x, &y)
		P[i] = NewPoint(i, x, y)
		T[i].l = NIL
		T[i].r = NIL
		T[i].p = NIL
	}
	fmt.Println(T)
	np = 0
	root := makeKDTree(0, N, 0)

	fmt.Println(P)
	fmt.Println(T)
	var q int
	fmt.Scanf("%d", &q)
	var sx, tx, sy, ty int
	for i := 0; i < q; i++ {
		fmt.Scanf("%d %d %d %d", &sx, &tx, &sy, &ty)
		ans := make(Ps, q)
		find(root, sx, tx, sy, ty, 0, ans)
		sort.Sort(ById{ans[:]})
		for j := 0; j < len(ans); j++ {
			ans[j].disp()
		}
	}
}
