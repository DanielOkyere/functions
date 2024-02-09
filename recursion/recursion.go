package recursion

import "fmt"

func main() {
	fact := factorial(9)
	fmt.Println(fact)
}

func factorial(number int) int {
	if number == 0{
		return 1
	}
	return number * factorial(number - 1)
}

// factorial of 5 => 5 * 4 * 3 * 2 * 1
