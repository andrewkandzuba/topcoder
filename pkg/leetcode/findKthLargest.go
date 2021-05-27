package leetcode

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	mergeSort(nums, 0, n-1)

	return nums[Max(n-k, 0)]
}

func mergeSort(A []int, l int, r int) {
	if l < r {
		var m = l + (r-l)/2
		mergeSort(A, l, m)
		mergeSort(A, m+1, r)
		mergeInternal(A, l, m, r)
	}
}

