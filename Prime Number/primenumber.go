package main

import "fmt"

func main() {

	number := 4

	result := checkprimeNumber(number)

	fmt.Print(result)

}

func checkprimeNumber(number int) string {

	if number <= 1 {
		return "Given number is not a prime number"
	}

	//iterates from 2 to the square root of the number, optimising the code
	for i := 2; i*i <= number; i++ {

		if number%i == 0 {
			return "Given number is a not prime number"
		}
	}

	return "Given number is a prime number"
}
