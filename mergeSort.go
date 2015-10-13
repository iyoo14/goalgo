package main

import (
	"fmt"
)

const (
	MAX      int = 500000
	SENTINEL int = 2000000000
)

var L [MAX/2 + 2]int
var R [MAX/2 + 2]int
var cnt int

func merge(A []int, n int, left int, mid int, right int) {
	n1 := mid - left
	n2 := right - mid
	for i := 0; i < n1; i++ {
		L[i] = A[left+i]
	}
	for i := 0; i < n2; i++ {
		R[i] = A[mid+i]
	}
	L[n1], R[n2] = SENTINEL, SENTINEL
	i := 0
	j := 0
	for k := left; k < right; k++ {
		cnt++
		if L[i] <= R[j] {
			A[k] = L[i]
			i++
		} else {
			A[k] = R[j]
			j++
		}
	}
}

func mergeSort(A []int, n int, left int, right int) {
	if left+1 < right {
		mid := (left + right) / 2
		mergeSort(A, n, left, mid)
		mergeSort(A, n, mid, right)
		merge(A, n, left, mid, right)
	}
}

func disp(A [MAX]int) {
	for i := 0; i < 9; i++ {
		fmt.Printf("%d ", A[i])
	}
	fmt.Printf("\n")
}

func main() {
	var A = [MAX]int{5, 2, 6, 3, 1, 9, 7, 8, 4}
	n := 9

	disp(A)
	mergeSort(A[:], n, 0, n)
	disp(A)
	fmt.Println(cnt)
}
