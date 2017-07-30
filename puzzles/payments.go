package puzzles

import "fmt"

func findComprise(incoming []int, target int) []int {
	comprised := make([]int, 0)

	fmt.Println(incoming)

	for i, v := range incoming {

		if v > target {
			continue
		}

		if v == target {
			comprised = append(comprised, v)
			continue
		}

		c := findComprise(incoming[i+1:], target-v)

		if len(c) > 0 {
			for _, cv:= range c {
				comprised = append(comprised, cv)
			}
		}
	}

	return comprised
}

func find(index int, incoming []int, comprised []int, target int) []int {
	if index < 0 || index == len(incoming) {
		if target == 0 {
			return comprised
		} else {
			return make([]int, 0)
		}
	}

	c, t := NewTarget(comprised, target, incoming[index]);
	lf := find(0, incoming[0:index-1], c, t)
	rf := find(index-1, incoming, c, t)

	if len(lf) > 0 {
		return lf
	}
	return rf
}

func NewTarget(comprised []int, target int, amount int) ([]int, int) {
	deduct := target - amount

	if deduct >= 0 {
		return append(comprised, amount), deduct
	}

	return comprised, target;
}
