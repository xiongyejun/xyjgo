package stack

import "errors"

type Stack struct {
	top int
	arr []interface{}
}

func (s *Stack) Len() int {
	return len(s.arr)
}

func (s *Stack) IsEmpty() bool {
	return s.top == 0
}

func (s *Stack) Push(x interface{}) {
	if len(s.arr) == s.top {
		s.arr = append(s.arr, x)
	} else {
		s.arr[s.top] = x
	}

	s.top++
	return
}

func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("empty stack")
	}
	s.top--
	return s.arr[s.top], nil
}

func (s *Stack) Top() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("empty stack")
	}
	return s.arr[s.top-1], nil
}

func New(maxNum int) *Stack {
	s := new(Stack)
	s.top = 0
	s.arr = make([]interface{}, 0, maxNum)
	return s
}

func (s *Stack) Do(f func(args ...interface{}) (int, error)) {
	f(s.arr[:s.top])
}
