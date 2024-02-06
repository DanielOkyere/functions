package functionsarevalues

import "fmt"

type transformfn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	morenumbers := []int{5, 1, 2}
	doubled := transformNumbers(&numbers, double)
	triple := transformNumbers(&numbers, triple)
	fmt.Println(doubled)
	fmt.Println(triple)

	tfFn1 := getTransformerfunction(&numbers)
	tfFn2 := getTransformerfunction(&morenumbers)

	transformedNumbers := transformNumbers(&numbers, tfFn1)
	moreTransformdedNumbers := transformNumbers(&morenumbers, tfFn2)

	fmt.Println(transformedNumbers)
	fmt.Println(moreTransformdedNumbers)
}

func getTransformerfunction(numbers *[]int) transformfn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func transformNumbers(numbers *[]int, transform transformfn) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
