package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MAX = 112
const SMAX = MAX * 2
const (
	EMPTY = iota
	ROOT
	NODE
)

var RAMDOMS [MAX]int
var SEARCHS [SMAX]int

type KEY int

func (k KEY) keyequal(n KEY) bool {
	if k == n {
		return true
	}
	return false
}

func (k KEY) keylt(n KEY) bool {
	if k < n {
		return true
	}
	return false
}

type node struct {
	stat  int
	data  KEY
	left  *node
	right *node
}

func (p *node) insert(key KEY) bool {
	stat := p.stat
	if stat == EMPTY {
		p.stat = ROOT
		p.data = key
		p.left = nil
		p.right = nil
		return true
	}
	var pp **node
	pp = &p
	for *pp != nil {
		if key.keyequal((*pp).data) {
			return false
		} else if key.keylt((*pp).data) {
			pp = &(*pp).left
		} else {
			pp = &(*pp).right
		}
	}
	n := new(node)
	n.stat = NODE
	n.data = key
	n.left = nil
	n.right = nil
	*pp = n
	return true
}

func (p *node) search(key KEY) *node {
	for p != nil {
		if key.keyequal(p.data) {
			return p
		} else if key.keylt(p.data) {
			p = p.left
		} else {
			p = p.right
		}
	}
	return nil
}

func traverse(n *node) {
	if n == nil {
		return
	}
	fmt.Println(n)
	traverse(n.left)
	traverse(n.right)
}

func pre() {
	rand.Seed(time.Now().UnixNano())

	nums := make([]int, MAX)
	nums = rand.Perm(MAX)
	for i, n := range nums {
		RAMDOMS[i] = n
	}
	fmt.Printf("RAMDOMS len is %d\n", len(RAMDOMS))

	snums := make([]int, SMAX)
	snums = rand.Perm(SMAX)
	for i, n := range snums {
		SEARCHS[i] = n
	}
	fmt.Printf("SEARCHS len is %d\n", len(SEARCHS))
}

func binaritree() {
	fmt.Printf("BINARI SEARCH\n")
	fmt.Printf("    start: insert.\n")
	start := time.Now()
	var tree *node
	tree = new(node)
	for _, i := range RAMDOMS {
		tree.insert(KEY(i))
	}
	end := time.Now()
	sub := end.Sub(start)
	fmt.Printf("    end  : insert time: %v.\n", sub)

	result := []int{}
	fmt.Printf("    start: search.\n")
	start = time.Now()
	for _, i := range SEARCHS {
		res := tree.search(KEY(i))
		if res != nil {
			result = append(result, int(res.data))
		}
	}
	end = time.Now()
	sub = end.Sub(start)
	fmt.Printf("    end  : search result num: %d, time: %v.\n", len(result), sub)
}

func golist() {
	fmt.Printf("GOLIST\n")
	fmt.Printf("    start: insert.\n")
	start := time.Now()
	var list [MAX]int
	for k, i := range RAMDOMS {
		list[k] = i
	}
	end := time.Now()
	sub := end.Sub(start)
	fmt.Printf("    end  : insert time: %v.\n", sub)

	result := []int{}
	fmt.Printf("    start: insert.\n")
	start = time.Now()
	for _, i := range SEARCHS {
		for _, v := range list {
			if i == v {
				result = append(result, i)
				break
			}
		}
	}
	end = time.Now()
	sub = end.Sub(start)
	fmt.Printf("    end  : search result num: %d, time: %v.\n", len(result), sub)
}

func main() {

	pre()
	binaritree()
	golist()
}
