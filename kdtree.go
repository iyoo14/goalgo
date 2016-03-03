package main

import (
	"fmt"
	"log"
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
	fmt.Printf("id - %d\n", p.id)
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

	//fmt.Printf("sx <= x : %d %d,  x <= tx : %d %d,sy <= y : %d %d,  y <= ty : %d  %d\n", sx, x, x, tx, sy, y, y, ty)
	if sx <= x && x <= tx && sy <= y && y <= ty {
		//fmt.Printf("%d %d %d %d %d %d\n", x, sx, tx, y, sy, ty)
		log.Println("append ", P[T[v].location])
		ans = append(ans, P[T[v].location])
		log.Println("len ", len(ans), " cap ", cap(ans))
		log.Println("appended ", ans)
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
				find(T[v].l, sx, tx, sy, ty, depth+1, ans)
			}
		}
		if T[v].r != NIL {
			if y <= ty {
				find(T[v].r, sx, tx, sy, ty, depth+1, ans)
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
	log.Println(T)
	np = 0
	root := makeKDTree(0, N, 0)

	log.Println(P)
	log.Println(T)
	var q int
	fmt.Scanf("%d", &q)
	var sx, tx, sy, ty int
	for i := 0; i < q; i++ {
		fmt.Scanf("%d %d %d %d", &sx, &tx, &sy, &ty)
		ans := make(Ps, 0, N)
		log.Println(ans)
		find(root, sx, tx, sy, ty, 0, ans)
		log.Println("ans ", ans[:cap(ans)])
		log.Println(len(ans), "  ", cap(ans))
		sort.Sort(ById{ans[:]})
		log.Println("sorted ans ", ans[:cap(ans)])
		for j := 0; j < cap(ans); j++ {
			log.Println(ans[j : j+1])
			an := ans[j : j+1]
			an[0].disp()
		}
	}
}
