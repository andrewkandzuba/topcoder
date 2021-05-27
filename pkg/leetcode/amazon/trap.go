package amazon

import "github.com/andrewkandzuba/topcoder/pkg/leetcode"

func trap(height []int) int {

	sp := 0
	ep := len(height) - 1
	maxV := 0

	for sp < ep {
		if height[sp] < height[ep] {
			i := sp
			for ; i < ep && height[i] <= height[sp]; i++ {
			}
			maxV = leetcode.Max(maxV, (ep-sp)*height[sp])
			sp = i
		} else {
			j := ep
			for ; j > sp && height[j] <= height[ep]; j-- {
			}
			maxV = leetcode.Max(maxV, (ep-sp)*height[ep])
			ep = j
		}
	}
	return maxV
}
