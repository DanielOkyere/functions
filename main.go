package main

import "fmt"

func main() {

	sum := sumup(1, 2, 4,5)
	fmt.Println(sum)
}

func sumup(numbers ...int) int {
	sum := 0
	for _, val := range numbers {
		sum = sum + val
	}

	return sum
}