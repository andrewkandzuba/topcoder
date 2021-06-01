package amazon

func minMeetingRooms(intervals [][]int) int {
	var root *BSTNode
	for i :=0; i < len(intervals); i++ {
		root = insertBST(root, intervals[i][0], 0)
		root = insertBST(root, intervals[i][1], 1)
	}
	var count, max int
	visitSortOrder(root, &count, &max)
	return max
}

func insertBST(node *BSTNode, key int, t byte) *BSTNode {
	if node == nil {
		return &BSTNode{
			key : key,
			t : t,
		}
	}
	if (key < node.key) || (key == node.key && t == 1) {
		node.left = insertBST(node.left, key, t)
	} else {
		node.right = insertBST(node.right, key, t)
	}
	return node
}

func visitSortOrder(node *BSTNode, count *int, max *int) {
	if node != nil {
		visitSortOrder(node.left, count, max);
		if node.t == 0 {
			*count++
		} else {
			*count--
		}
		if *count > *max {
			*max = *count
		}
		println("Node:", node.key, node.t, *count)
		visitSortOrder(node.right, count, max);
	}
}

type BSTNode struct {
	key int
	t byte // 0 - begin, 1 - end
	left, right *BSTNode
}


