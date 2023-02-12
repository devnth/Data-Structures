package main

// func main() {

// 	data := []int{2, 4, 5, 6, 7, 4, 3}

// 	fmt.Printf("Index of Target %d", ans(data))

// }

func answer(data []int) int {

	start := 0
	end := len(data) - 1

	for start < end {

		mid := start + (end-start)/2

		if data[mid] < data[mid+1] {
			start = mid + 1
		} else {
			end = mid
		}

	}
	return data[start]
}
