package trie

const(
	MAX_CHAR = 26
	START_CHAR = 'a'
)


type Index struct {
	root *Node
}

type Node struct {
	children []*Node
	isLeaf   bool
	score    int
}

func NewIndex() *Index {
	return &Index{
		root: &Node{
			children: make([]*Node, MAX_CHAR),
			isLeaf:   false,
			score:    0,
		},
	}
}

func (index *Index) Insert(str string) {
	children := index.root.children
	for i := 0; i < len(str); i++ {
		c := str[i] - START_CHAR
		n := children[c]
		if n == nil {
			n = &Node{
				children: make([]*Node, MAX_CHAR),
				isLeaf:   false,
				score:    0,
			}
			children[c] = n
		}
		children = n.children
		if i == len(str) - 1 {
			n.isLeaf = true
		}
	}
}

func (index *Index) Search(word string) bool {
	n := index.findNode(word)
	if n == nil {
		return false
	}
	return n.isLeaf
}

func (index *Index) SearchPrefix(prefix string) bool {
	n := index.findNode(prefix)
	return n != nil
}

func (index *Index) findNode(str string) *Node {
	t := index.root
	for i:=0; i < len(str); i++ {
		c := str[i] - START_CHAR
		if t.children[c] != nil {
			t = t.children[c]
		} else {
			return nil
		}
	}
	if t == index.root {
		return nil
	}
	return t
}


