package main

import "fmt"

func main() {

	array := []int{7, 2, 5, 8, 10}

	result := splitArray(array, 2)

	fmt.Println(result)
}

// function to find the largest sum from the splited array
func splitArray(nums []int, k int) int {

	start := 0
	end := 0

	max := nums[0]
	for _, v := range nums {

		// to find the largest number in an array
		// which is the smallest possible largest sum
		if max < v {
			max = v
		}

		// to find the largest possible sum
		// which is the sum of all elements in an array
		end = end + v

	}

	start = max
	// fmt.Printf("Min Range %v, Max Range %v\n", start, end)

	//to find a number between a range we use binary search

	for start < end {

		mid := start + (end-start)/2

		sum := 0
		pieces := 1

		for _, num := range nums {

			if sum+num > mid {
				sum = num
				pieces++
			} else {
				sum = sum + num
			}
		}

		if pieces > k {

			start = mid + 1 // pieces are more than required
			// so mid should be larger

		} else {
			end = mid
		}

	}

	return end //here start = end

}
