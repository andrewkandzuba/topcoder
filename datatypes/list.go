package datatypes

type ListNode interface {
	next() ListNode
}

type List interface {
	add(n ListNode)
}
