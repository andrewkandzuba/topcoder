package leetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1, nums2 = sortByLength(nums1, nums2)

	if len(nums1) == 0 {
		m, _, _ := median(nums2)
		return m
	}

	if len(nums2) == 0 {
		m, _, _ := median(nums1)
		return m
	}

	if len(nums1) == 1 && len(nums2) == 1 {
		return float64(nums1[0]+nums2[0]) / 2
	}
	if len(nums1) == 1 {
		return medianSingleAndMultiple(nums1, nums2)
	}
	if len(nums1) == 2 {
		return medianTwoAndMultiple(nums1, nums2)
	}

	for len(nums1) > 2 {
		m1, l1, r1 := median(nums1)
		m2, l2, r2 := median(nums2)
		if m1 < m2 {
			p := min(l1, len(nums2)-r2-1)
			nums1 = nums1[p:]
			nums2 = nums2[:len(nums2)-p]
		} else {
			p := min(len(nums1)-r1-1, l2)
			nums1 = nums1[:len(nums1)-p]
			nums2 = nums2[p:]
		}
	}

	return medianTwoAndMultiple(nums1, nums2)
}

func sortByLength(nums1 []int, nums2 []int) ([]int, []int) {
	if len(nums1) > len(nums2) {
		return nums2, nums1
	}
	return nums1, nums2
}

func medianSingleAndTwo(single []int, nums []int) float64 {
	if single[0] < nums[0] {
		return float64(nums[0])
	}
	if single[0] > nums[1] {
		return float64(nums[1])
	}
	return float64(single[0])
}

func medianTwoAndTwo(nums1 []int, nums2 []int) float64 {
	m, _, _ := median([]int{min(nums1[0], nums2[0]), max(nums1[0], nums2[0]), min(nums1[1], nums2[1]), max(nums1[1], nums2[1])})
	return m
}

func medianSingleAndMultiple(single []int, nums []int) float64 {
	var n = len(nums)
	if len(nums) == 2 {
		return medianSingleAndTwo(single, nums)
	}
	if n%2 == 0 {
		if single[0] < nums[n/2-1] {
			return float64(nums[n/2-1])
		}
		if single[0] >= nums[n/2-1] && single[0] <= nums[n/2] {
			return float64(single[0])
		}
		return float64(nums[n/2])
	}
	if single[0] < nums[n/2-1] {
		return float64(nums[n/2-1]+nums[n/2]) / 2.0
	}
	if single[0] >= nums[n/2-1] && single[0] <= nums[n/2+1] {
		return float64(single[0]+nums[n/2]) / 2.0
	}
	return float64(nums[n/2]+nums[n/2+1]) / 2.0
}

func medianTwoAndMultiple(two []int, nums []int) float64 {
	var n = len(nums)
	if n == 2 {
		return medianTwoAndTwo(two, nums)
	}
	if n%2 == 0 {
		m, _, _ := median(merge(two, []int{nums[n/2-2], nums[n/2-1], nums[n/2], nums[n/2+1]}))
		return m
	}
	if n == 3 {
		m, _, _ := median(merge(two, []int{nums[0], nums[1], nums[2]}))
		return m
	}
	m, _, _ := median(merge(two, []int{nums[n/2-2], nums[n/2-1], nums[n/2], nums[n/2+1], nums[n/2+1]}))
	return m
}

func median(nums []int) (float64, int, int) {
	var n = len(nums)
	if len(nums)%2 == 0 {
		l := n/2 - 1
		r := n / 2
		return float64(nums[l]+nums[r]) / 2.0, l, r
	}
	l := n / 2
	r := n / 2
	return float64(nums[l]), l, r
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func merge(a []int, b []int) []int {
	var r = make([]int, 0)
	i := 0
	j := 0
	for i < len(a) {
		for j < len(b) && a[i] > b[j] {
			r = append(r, b[j])
			j++
		}
		r = append(r, a[i])
		i++
	}
	for j < len(b) {
		r = append(r, b[j])
		j++
	}
	return r
}
