package occ

import (
	"fmt"
	"github.com/andrewkandzuba/topcoder/pkg/datatypes/trie"
)

func PrintSorted(arr []string) {
	index := trie.NewIndex()
	for i := 0; i < len(arr); i++ {
		index.Insert(arr[i], i)
	}
	preorder(index.Root, arr)
}

func preorder(node *trie.Node, arr []string) {
	if node == nil {
		return
	}
	for i := 0; i < trie.MaxChar; i++ {
		if node.Children[i] != nil {
			if node.Children[i].Pos != -1 {
				fmt.Println(arr[node.Children[i].Pos])
			}
			preorder(node.Children[i], arr)
		}
	}
}

func Score(arr []string) int {
	index := trie.NewIndex()
	for i := 0; i < len(arr); i++ {
		index.Insert(arr[i], i)
	}
	score := 0
	order := 1
	preorderScore(index.Root, arr, &score, &order)
	return score
}

func preorderScore(node *trie.Node, arr []string, score *int, order *int) {
	if node == nil {
		return
	}
	for i := 0; i < trie.MaxChar; i++ {
		if node.Children[i] != nil {
			if node.Children[i].Pos != -1 {
				*score += node.Children[i].Score * *order
				*order++
			}
			preorderScore(node.Children[i], arr, score, order)
		}
	}
}
