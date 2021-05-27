package amazon

func trap(height []int) int {
	totalTrapped := 0

	sp := 0
	ep := len(height) - 1

	for sp < ep {
		projected := 0
		if height[sp] < height[ep] {
			lp := sp + 1
			for ep > lp && height[lp] < height[sp] {
				projected += height[sp] - height[lp]
				lp++
			}
			sp = lp
		} else {
			rp := ep - 1
			for sp < rp && height[ep] > height[rp] {
				projected += height[ep] - height[rp]
				rp--
			}
			ep = rp
		}
		totalTrapped += projected
	}
	return totalTrapped
}