package sort

func MergeSort(A []int) []int {
	var B = make([]int, len(A))
	copy(A, B)
	split(B, 0, len(A), A)
	return A
}

func split(A []int, iBegin int, iEnd int, B []int) {
	if iEnd-iBegin < 2 {
		return
	}
	var iMid = (iEnd + iBegin) / 2
	split(A, iBegin, iMid, B)
	split(A, iMid, iEnd, B)
	merge(B, iBegin, iMid, iEnd, A)
}

func merge(A []int, iBegin int, iMid int, iEnd int, B []int) {
	var i = iBegin
	var j = iMid

	for k := iBegin; k < iEnd; k++{
		if i < iMid && (j >= iEnd || A[i] <= A[j]) {
			B[k] = A[i]
			i = i + 1
		} else {
			B[k] = A[j]
			j = j + 1
		}
	}
}
