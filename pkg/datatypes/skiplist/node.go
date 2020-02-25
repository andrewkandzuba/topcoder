package skiplist

type Node struct {
	key int
	forward []*Node
}

func NewNode(key int, level uint) *Node {
	return &Node{
		key: key,
		forward: make([]*Node, level + 1),
	}
}