package leetcode

func maxArea(height []int) int {

	sp := 0
	ep := len(height) - 1
	n := ep - sp
	maxV := n * min(height[sp], height[ep])

	i := sp
	j := ep
	for i < j {
		outerV := 2 * min(height[sp], height[i+1])
		rightV := min(height[sp+1], height[i+2])
		if rightV > outerV {
			sp++
		}

		outerV = 2 * min(height[ep-1], height[ep])
		leftV := min(height[ep-2], height[ep-1])
		if leftV > outerV {
			ep--
		}
		maxV = max(maxV, maxArea(height[sp:ep+1]))
	}

	return maxV
}
