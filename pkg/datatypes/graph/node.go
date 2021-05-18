package graph

type Node struct {
	id int
	children []*Node
}

func NewNode(id int) *Node {
	return &Node{id, make([]*Node, 0)}
}

func (node *Node) AddChild(child *Node) {
	node.children = append(node.children, child)
}