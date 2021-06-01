package amazon

func minMeetingRooms2(intervals [][]int) int {
	var root *ITNode = nil
	for i :=0; i < len(intervals); i++ {
		root = insert(root, Interval{intervals[i][0], intervals[i][1]})
	}
	return 0
}

func insert(root *ITNode, i Interval) *ITNode {
	if root == nil {
		return &ITNode{
			i:     &i,
			max:   i.high,
			left:  nil,
			right: nil,
		}
	}

	l := root.i.low

	if i.low < l {
		root.left = insert(root.left, i)
	} else {
		root.right = insert(root.right, i)
	}

	if root.max < i.high {
		root.max = i.high
	}

	return root
}

func isOverlapped(i1 *Interval, i2 *Interval) bool {
	if i1.low <= i2.high && i2.low <= i1.high {
		return true
	}
	return false
}

type Interval struct {
	low int
	high int
}

type ITNode struct {
	i *Interval
	max int
	left, right *ITNode
}


