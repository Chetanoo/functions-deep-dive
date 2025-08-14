package main

import "fmt"

type transformerFunction func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	doubled := transformNumbers(&numbers, double)
	fmt.Println(doubled)

	tripled := transformNumbers(&numbers, triple)
	fmt.Println(tripled)
}

func transformNumbers(numbers *[]int, transformer transformerFunction) []int {
	dNumbers := make([]int, len(*numbers))
	for index, number := range *numbers {
		dNumbers[index] = transformer(number)
	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
