package main

import (
	"fmt"
)

func main() {

	characters := []byte{'c', 'f', 'j'}

	target := byte('d')

	result := nextSmallestCharacter(characters, target)

	fmt.Println("output:  ", result)

}

func nextSmallestCharacter(characters []byte, target byte) string {

	start := 0

	end := len(characters) - 1

	for start <= end {

		middle := start + (end-start)/2

		if characters[middle] == target {

			if middle == len(characters)-1 {
				return string(characters[0])
			}
			return string(characters[middle+1])
		} else if target <= characters[middle] {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}

	return string(characters[start])

}
