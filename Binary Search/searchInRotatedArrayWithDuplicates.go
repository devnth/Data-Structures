package main

func findPivot(data []int) int {

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

		//to check whether if duplicates are present at the start or end
		//if duplicates are present at the start or end then move the start to next index and
		// move end to the previous index
		if data[mid] == data[start] && data[mid] == data[end] {

			//check whether start is pivot
			if data[start] > data[start+1] {
				return start
			}
			start++

			//check whether end is pivot
			if data[end] < data[end-1] {
				return end
			}
			end--
		} else if data[mid] > data[start] || data[mid] == data[start] && data[mid] > data[end] {
			start = mid + 1
		} else {
			end = mid - 1
		}

	}
	return -1
}
