package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	pr := &ListNode{0, nil}

	p1 := l1
	p2 := l2
	pc := pr
	r := 0

	for p1 != nil || p2 != nil {
		s := r
		if p1 != nil {
			s = s + p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			s = s + p2.Val
			p2 = p2.Next
		}

		r = s / 10

		pc.Val = s % 10
		if r > 0 {
			pc.Next = &ListNode{0, nil}
			pc = pc.Next
		}
	}

	if r > 0 {
		pc.Next = &ListNode{r, nil}
		pc = pc.Next
	}

	return pr
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if j != i && nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
