package main

// find an element in an infinite sorted array you have to find the range
// first. One possible way is to increase the size of the range exponentially
// time complexity for finding the range is log(N)
func findRange(array []int, target int) (int, int) {

	start := 0
	end := 1

	for target > array[end] {

		newstart := end + 1

		//end = end + size of the array multplied by 2
		end = end + (end-start+1)*2
		start = newstart

	}

	return start, end

}

// after getting the range, do normal binary search
func binarySearchInfiniteArray(array []int, target int, start int, end int) int {

	for start <= end {
		mid := start + (end-start)/2

		if target == array[mid] {
			return mid
		} else if target > array[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1

}
