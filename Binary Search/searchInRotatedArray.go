package main

func search(data []int, target int) int {

	pivot := findpivot(data)

	if pivot == -1 {
		return binarysearch(data, target, 0, len(data)-1)
	}

	if data[pivot] == target {
		return pivot
	}
	if data[0] > target {
		return binarysearch(data, target, pivot+1, len(data)-1)
	}

	return binarysearch(data, target, 0, pivot-1)

}

func findpivot(data []int) int {

	start := 0
	end := len(data) - 1

	for start <= end {

		mid := start + (end-start)/2

		if mid > start && data[mid] < data[mid-1] {
			return mid
		}
		if mid < end && data[mid] > data[mid+1] {
			return mid
		}

		if data[start] > data[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}

	}
	return -1
}

func binarysearch(data []int, target int, start int, end int) int {

	for start <= end {

		mid := start + (end-start)/2

		if target < data[mid] {
			end = mid - 1
		} else if target > data[mid] {
			start = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
