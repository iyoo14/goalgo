package lib

type Stack struct {
	top int
	sk  []int
	max int
}

type MyError struct {
	msg string
}

func (err MyError) Error() string {
	return err.msg
}

func NewStack(sz int) *Stack {
	s := new(Stack)
	s.top = 0
	s.max = sz
	s.sk = make([]int, sz)
	return s
}

func (s *Stack) IsEmpty() bool {
	return s.top == 0
}

func (s *Stack) IsFull() bool {
	return s.top >= s.max-1
}

func (s *Stack) Push(x int) (err error) {
	if s.IsFull() {
		return MyError{"is full"}
	}
	s.top++
	s.sk[s.top] = x
	return
}

func (s *Stack) Pop() (n int, err error) {
	if s.IsEmpty() {
		return 0, MyError{"under flow"}
	}
	s.top--
	n = s.sk[s.top+1]
	return
}
