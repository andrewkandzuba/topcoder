package graph

import "testing"

func TestDfs_Traverse(t *testing.T) {
	n1 := NewNode(1) // A
	n2 := NewNode(2) // B
	n3 := NewNode(3) // C
	n4 := NewNode(4) // E
	n5 := NewNode(5) // D

	n1.AddChild(n2)
	n1.AddChild(n5)
	n1.AddChild(n4)

	n2.AddChild(n3)
	n2.AddChild(n5)

	n5.AddChild(n1)
	n5.AddChild(n4)
}