package trie

const(
	MaxChar   = 26
	StartChar = 'A'
)

type Node struct {
	Children []*Node
	Pos      int
	Score    int
}

func newNode() *Node {
	return &Node{
		Children: make([]*Node, MaxChar),
		Pos:      -1,
	}
}

type Index struct {
	Root *Node
}

func NewIndex() *Index {
	return &Index{newNode()}
}

func (index *Index) Insert(str string, pos int) {
	children := index.Root.Children
	totalScore := 0
	for i := 0; i < len(str); i++ {
		c := str[i] - StartChar
		totalScore += int(c) + 1
		n := children[c]
		if n == nil {
			n = newNode()
			children[c] = n
		}
		children = n.Children
		if i == len(str) - 1 {
			if n.Pos == -1 {
				n.Pos = pos
				n.Score = totalScore
			}
		}
	}
}

func (index *Index) Search(word string) bool {
	n := index.findNode(word)
	if n == nil {
		return false
	}
	return n.Pos != -1
}

func (index *Index) SearchPrefix(prefix string) bool {
	n := index.findNode(prefix)
	return n != nil
}

func (index *Index) findNode(str string) *Node {
	t := index.Root
	for i:=0; i < len(str); i++ {
		c := str[i] - StartChar
		if t.Children[c] != nil {
			t = t.Children[c]
		} else {
			return nil
		}
	}
	if t == index.Root {
		return nil
	}
	return t
}

