package tree

type Node struct {
	Value        int
	Left, Rright *Node
}

func NewNode(value int) *Node {
	return &Node{Value: value}
}

// 遍历
func (me *Node) Traversal(f func(*Node)) {
	if me == nil {
		return
	} else {
		me.Left.Traversal(f)
		f(me)
		me.Rright.Traversal(f)
	}
}
