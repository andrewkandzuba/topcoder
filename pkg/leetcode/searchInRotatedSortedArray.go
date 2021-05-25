package leetcode

func searchInArray(nums[] int, target int) int {
	n := len(nums)
	pivot := findPivot(nums, 0, n - 1)

	if pivot == -1 {
		return binarySearch(nums, 0, n - 1, target)
	}

	if nums[pivot] == target {
		return pivot;
	}

	if nums[0] <= target {
		return binarySearch(nums, 0, pivot-1, target);
	}

	return binarySearch(nums, pivot + 1, n - 1, target);
}

func binarySearch(arr[] int, low int, high int, key int) int {
	if high < low {
		return -1;
	}

	mid := (low + high) / 2
	if key == arr[mid] {
		return mid;
	}

	if key > arr[mid] {
		return binarySearch(arr, mid + 1, high, key)
	}

	return binarySearch(arr, low, mid - 1, key);
}

func findPivot(arr[] int, low int, high int) int {
	if high < low {
		return -1;
	}
	if high == low {
		return low;
	}

	mid := (low + high) / 2
	if mid < high && arr[mid] > arr[mid + 1] {
		return mid;
	}

	if mid > low && arr[mid] < arr[mid - 1] {
		return mid - 1;
	}

	if arr[low] >= arr[mid] {
		return findPivot(arr, low, mid-1)
	}

	return findPivot(arr, mid + 1, high)
}