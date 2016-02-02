package lib

type Vector struct {
	V   []int
	idx int
}

func NewVector(n int) Vector {
	vec := make([]int, n)
	return Vector{vec, 0}
}

func (v *Vector) PushBack(val int) {
	v.V[v.idx] = val
	v.idx++
}

func (v *Vector) Size() int {
	return len(v.V)
}
