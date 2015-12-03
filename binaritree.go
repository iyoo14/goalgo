package main

import (
	"fmt"
)

const (
	EMPTY = iota
	ROOT
	NODE
)

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

func main() {

	var root *node
	root = new(node)
	root.insert(12)
	root.insert(4)
	root.insert(8)
	root.insert(2)
	root.insert(6)

	fmt.Println("main")
	traverse(root)
	res := root.search(2)
	fmt.Println(res)
}
