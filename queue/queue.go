package queue

// an FIFO queue
type Queue []interface{}

// Pushes the element into the queue.
func (me *Queue) Push(v interface{}) {
	*me = append(*me, v)
}

// Pops element from head.
func (me *Queue) Pop() interface{} {
	if me.IsEmpty() {
		return nil
	}
	head := (*me)[0]
	*me = (*me)[1:]
	return head
}

// Return queue is empty
func (me *Queue) IsEmpty() bool {
	return len(*me) == 0
}
