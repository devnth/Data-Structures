package main

// func main() {

// 	characters := []byte{'c', 'f', 'j'}

// 	target := byte('d')

// 	result := nextSmallestCharacter(characters, target)

// 	fmt.Println("output:  ", result)

// }

func nextSmallestCharacter(characters []byte, target byte) string {

	start := 0

	length := len(characters)

	end := length - 1

	for start <= end {

		middle := start + (end-start)/2

		if target < characters[middle] {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}

	return string(characters[start%length])
	//we are taking the remainder so that if next smallest number doesn't exits
	//we return the first letter in the array

}
