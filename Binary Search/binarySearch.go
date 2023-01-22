package main

import "fmt"

func main() {

	array := []int{3, 5, 9, 23, 30, 34, 45, 56, 78}

	target := 2

	status, index := binarySearch(array, target)

	if status {
		fmt.Printf("target found at %d", index)
	} else {
		fmt.Println("target not found")
	}

}

func binarySearch(array []int, target int) (bool, int) {

	start := 0
	end := len(array) - 1

	for start <= end {

		middle := start + (end-start)/2

		if array[middle] == target {
			return true, middle + 1
		} else if array[middle] > target {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}

	return false, 0
}
