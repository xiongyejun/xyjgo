package queue

import (
	"testing"
)

func Test_func(t *testing.T) {
	q := new(Queue)

	q.Push(1)
	q.Push(2)
	q.Push(3)
	t.Log(q.IsEmpty())
	q.Push("abb")
	t.Log(q.Pop())
	t.Log(q.Pop())
	t.Log(q.Pop())
	t.Log(q.Pop())
}
