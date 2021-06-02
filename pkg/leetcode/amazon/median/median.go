package median

import "math/rand"

type MedianFinder struct {
	arr []int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		arr: make([]int, 0),
	}
}


func (this *MedianFinder) AddNum(num int)  {
	this.arr = append(this.arr, num)
}


func (this *MedianFinder) FindMedian() float64 {
	var ans = 0.0
	var a = -1
	var b = -1
	var n = len(this.arr)

	if n % 2 == 1 {
		median(this.arr, 0, n - 1, n / 2, &a, &b)
		ans = float64(b)
	} else {
		median(this.arr, 0, n - 1, n / 2, &a, &b);
		ans = float64(a + b) / 2;
	}

	return ans
}

func median(arr []int, l int, r int, k int, a *int, b *int) {
	if l < r {
		pi := partition_r(arr, l, r)

		if pi == k {
			*b = arr[pi];
			if *a != -1 {
				return;
			}
		}  else if pi == k - 1 {
			*a = arr[pi];
			if *b != -1 {
				return;
			}
		}
		if pi >= k {
			median(arr, l,  pi-1, k, a, b)
		} else {
			median(arr, pi+1, r, k, a, b)
		}
	}
}

func partition(arr []int, lo int, hi int) int {
	pivot := arr[hi]
	i := lo     // place for swapping
	for j := lo; j < hi-1; j++ {
		if arr[j] <= pivot {
			swap(arr, i, j)
			i = i + 1
		}
	}
	swap(arr, i, hi)
	return i
}

func partition_r(arr []int, lo int, hi int) int {
	r := rand.Intn(hi - lo) + lo
	swap(arr, r, hi)
	return partition(arr, lo, hi)
}

func swap(arr []int, i int, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

/**
Randomly pick pivot element from arr[] and the using the partition step from the quick sort algorithm arrange all the elements smaller than the pivot on its left and the elements greater than it on its right.
If after the previous step, the position of the chosen pivot is the middle of the array then it is the required median of the given array.
If the position is before the middle element then repeat the step for the subarray starting from previous starting index and the chosen pivot as the ending index.
If the position is after the middle element then repeat the step for the subarray starting from the chosen pivot and ending at the previous ending index.
Note that in case of even number of elements, the middle two elements have to be found and their average will be the median of the array.

 */
/*
quicksort(arr[], lo, hi)
if lo < hi
p = partition_r(arr, lo, hi)
quicksort(arr, lo , p-1)
quicksort(arr, p+1, hi)
*/

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
