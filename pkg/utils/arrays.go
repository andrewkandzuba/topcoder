package utils

import "github.com/andrewkandzuba/topcoder/pkg/sort"

func SortAndCompress(A []int) []int {
	if len(A) < 2 {
		return A
	}

	A = sort.MergeSort(A)

	var B = append(make([]int,0), A[0])
	for i := 1; i < len(A); i++ {
		if B[len(B)-1] != A[i]{
			B = append(B, A[i])
		}
	}
	return B
}
