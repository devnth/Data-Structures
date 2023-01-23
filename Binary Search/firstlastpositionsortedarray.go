package main

import (
	"fmt"
)

func main() {

	data := []int{5, 7, 7, 7, 7, 8, 8, 10}
	target := 8

	start := searchTarget(data, target, true)
	end := searchTarget(data, target, false)

	fmt.Printf("Start Index: %d\nLast Index: %d", start, end)

}

func searchTarget(data []int, target int, startIndex bool) int {

	ans := -1

	start := 0
	end := len(data) - 1

	for start <= end {

		mid := start + (end-start)/2

		if target < data[mid] {
			end = mid - 1
		} else if target > data[mid] {
			start = mid + 1
		} else {

			ans = mid

			if startIndex {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
	}

	return ans + 1
}
