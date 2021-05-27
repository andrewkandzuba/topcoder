package amazon

func ladderLength(beginWord string, endWord string, wordList []string) int {
	dict := stringListToSet(wordList)
	if !dict[endWord] {
		return 0
	}

	var queue []wordNode
	queue = append(queue, wordNode{beginWord, 1})
	dict[beginWord] = false

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for k, v := range dict {
			if !v {
				continue
			}
			if diffInOneChar(cur.word, k) {
				if k == endWord {
					return cur.pos + 1
				}
				queue = append(queue, wordNode{k, cur.pos + 1})
				dict[k] = false
			}
		}
	}
	return 0
}

func stringListToSet(wordList []string) map[string]bool {
	s := make(map[string]bool)
	for _, w := range wordList {
		s[w] = true
    }
	return s
}

func diffInOneChar(source string, target string) bool {
	if len(source) != len(target) {
		return false
	}
	diff := 0
	n := len(source)
	for i := 0; i < n; i++ {
		if source[i] != target[i] {
			diff++
		}
		if diff > 1 {
			return false
		}
	}
	return diff == 1
}

type wordNode struct {
	word string
	pos int
}