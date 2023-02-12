package main

// func main() {

// 	data := []int{5, 7, 7, 7, 7, 8, 8, 10}
// 	target := 8

// 	start := searchTarget(data, target, true)
// 	end := searchTarget(data, target, false)

// 	fmt.Printf("Start Index: %d\nLast Index: %d", start, end)

// }

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

// func main() {

// 	data := []int{3,5,7,9,10,90,100,130,140,160,170}
// 	target := 80

// 	fmt.Printf("Index of Target %d",ans(data, target))

// }

func ans(data []int, target int) int {

	start := 0
	end := 1

	for target > data[end] {

		temp := end + 1
		end = end + (end-start+1)*2
		start = temp
	}

	// return binarySearch(data, target, start, end)
	return 0
}

// func binarySearch(data []int, target int, start int, end int) int {

// 	for start <= end {

// 		mid := start + (end-start)/2

// 		if target < data[mid] {
// 			end = mid - 1
// 		} else if target > data[mid] {
// 			start = mid + 1
// 		} else {
// 			return mid + 1
// 		}
// 	}

// 	return -1
// }
