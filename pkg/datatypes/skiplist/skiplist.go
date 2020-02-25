package skiplist

import (
	"fmt"
	"math/rand"
)

type SkipList struct {
	maxLvl uint
	level  uint
	p      float64
	header *Node
}

func NewSkipList(maxLvl uint, p float64) *SkipList {
	return &SkipList{
		maxLvl: maxLvl,
		level:  0,
		p:      p,
		header: NewNode(-1, maxLvl),
	}
}

func (sk *SkipList) Insert(key int) {

	var update = make([]*Node, sk.maxLvl+1)

	var current = sk.header
	for i := int(sk.level); i >= 0; i-- {
		for current.forward[i] != nil && current.forward[i].key < key {
			current = current.forward[i]
		}
		update[i] = current
	}

	current = current.forward[0]

	if current == nil || current.key != key {

		rLevel := sk.randomLevel()

		if rLevel > sk.level {
			for i := sk.level + 1; i < rLevel+1; i++ {
				update[i] = sk.header
				sk.level = rLevel
			}
		}

		n := NewNode(key, rLevel)

		for i := uint(0); i <= rLevel; i++ {
			n.forward[i] = update[i].forward[i]
			update[i].forward[i] = n
		}
	}
}

func (sk *SkipList) Contains(key int) bool {
	var current = sk.header
	for i := int(sk.level); i >= 0; i-- {
		for current.forward[i] != nil && current.forward[i].key < key {
			current = current.forward[i]
		}
	}
	current = current.forward[0]

	if current == nil || current.key != key {
		return false
	}

	return true
}

func (sk *SkipList) Print() {
	for i := uint(0); i <= sk.level; i++ {
		fmt.Printf("Level %d: ", i)
		var n = sk.header.forward[i]
		for n != nil {
			fmt.Printf("%d->", n.key)
			n = n.forward[i]
		}
		fmt.Printf("nil \n")
	}
}

func (sk *SkipList) randomLevel() uint {
	var lvl = uint(0)
	for rand.Float64() < sk.p && lvl < sk.maxLvl {
		lvl += 1
	}
	return lvl
}
