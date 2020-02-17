package skiplist

type SkipList struct {
	maxLevel int
	level int
	p float64
	header *Node
}

func NewSkipList() *SkipList {
	return nil
}