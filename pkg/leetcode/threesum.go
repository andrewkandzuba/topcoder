package leetcode

import "github.com/andrewkandzuba/topcoder/pkg/sort"

func threeSum(nums []int) [][]int {
	var slices = make([][]int, 0)
	var n = len(nums)
	if n < 3 {
		return slices
	}
	var A = sort.MergeSort(nums)
	i := 0
	for i < n-2 {
		l := i + 1
		r := n - 1
		for l < r {
			if A[i]+A[l]+A[r] == 0 {
				slice := make([]int, 0)
				slice = append(slice, A[i], A[l], A[r])
				slices = append(slices, slice)
				rm := A[r]
				r--
				for l < r && rm == A[r] {
					r--
				}
			} else if A[i]+A[l]+A[r] < 0 {
				rm := A[l]
				l++
				for l < r && rm == A[l] {
					l++
				}
			} else {
				rm := A[r]
				r--
				for l < r && rm == A[r] {
					r--
				}
			}
		}
		rm := A[i]
		i++
		for i < n-2 && rm == A[i] {
			i++
		}
	}

	return slices
}

func twoSum2(nums []int) [][]int {
	var slices = make([][]int, 0)

	var A = sort.MergeSort(nums)
	var n = len(A)
	var i = 0
	for i < n {
		var lm = A[i]
		j := i + 1
		for j < n {
			rm := A[j]
			if rm == -lm {
				slice := make([]int, 0)
				slice = append(slice, lm)
				slice = append(slice, rm)
				slices = append(slices, slice)
			}
			for j < n && rm == A[j] {
				j++
			}
		}
		for i < n && lm == A[i] {
			i++
		}
	}
	return slices
}
