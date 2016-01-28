package lib

type Queue struct {
	head int
	tail int
	max  int
	que  []interface{}
}

func NewQueue(sz int) *Queue {
	q := new(Queue)
	q.head = 0
	q.tail = 0
	q.max = sz
	q.que = make([]interface{}, sz)
	return q
}

func (q *Queue) IsEmpty() bool {
	return q.head == q.tail
}

func (q *Queue) IsFull() bool {
	return q.head == (q.tail+1)%q.max
}

func (q *Queue) Front() interface{} {
	return q.que[q.head]
}

func (q *Queue) Enqueue(x interface{}) (err error) {
	if q.IsFull() {
		return MyError{"is full"}
	}
	q.que[q.tail] = x
	if q.tail+1 == q.max {
		q.tail = 0
	} else {
		q.tail++
	}
	return
}

func (q *Queue) Dequeue() (n interface{}, err error) {
	if q.IsEmpty() {
		return 0, MyError{"under flow"}
	}
	n = q.que[q.head]
	if q.head+1 == q.max {
		q.head = 0
	} else {
		q.head++
	}
	return
}
