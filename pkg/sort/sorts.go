package sort

func MergeSort(A []int) []int {
	mergeSort(A, 0, len(A)-1)
	return A
}

func mergeSort(A []int, l int, r int) {
	if l < r {
		var m = l + (r-l)/2
		mergeSort(A, l, m)
		mergeSort(A, m+1, r)
		merge(A, l, m, r)
	}
}

func merge(A []int, l, m, r int) {
	var L = make([]int, m-l+1)
	var R = make([]int, r-m)

	var i = 0
	var j = 0
	var k = l

	copy(L, A[l:m+1])
	copy(R, A[m+1:r+1])

	for i < len(L) && j < len(R) {
		if L[i] < R[j] {
			A[k] = L[i]
			i++
		} else {
			A[k] = R[j]
			j++
		}
		k++
	}

	for i < len(L) {
		A[k] = L[i]
		i++
		k++
	}
	for j < len(R) {
		A[k] = R[j]
		j++
		k++
	}
}
