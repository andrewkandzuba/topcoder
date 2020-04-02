package leetcode

func threeSum(nums []int) [][]int {
	var slices = make([][]int, 0)
	for i, a := range nums {
		sumOfTwo(nums, i, - a, &slices)
	}
	return slices
}

func oneSum(nums []int, sum int, slices *[][]int) {
	for _, a := range nums {
		if a == sum {
			slice := make([]int, 0)
			slice = append(slice, a)
			*slices = append(*slices, slice)
		}
	}
}

func sumOfTwo(nums []int, idx int, sum int, slices *[][]int) {
	for i, a := range nums {
		if i > idx && a == sum {
			slice := make([]int, 0)
			slice = append(slice, nums[idx])
			slice = append(slice, a)
			*slices = append(*slices, slice)
		}
	}
}
