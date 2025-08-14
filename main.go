package main

import "fmt"

type transformerFunction func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	moreNumbers := []int{5, 6, 7, 8, 9}
	doubled := transformNumbers(&numbers, double)
	fmt.Println(doubled)

	tripled := transformNumbers(&numbers, triple)
	fmt.Println(tripled)

	transformer1 := getTransFormer(&numbers)
	transformer2 := getTransFormer(&moreNumbers)
	test1 := transformNumbers(&numbers, transformer1)
	test2 := transformNumbers(&moreNumbers, transformer2)
	fmt.Println(test1)
	fmt.Println(test2)
}

func transformNumbers(numbers *[]int, transformer transformerFunction) []int {
	dNumbers := make([]int, len(*numbers))
	for index, number := range *numbers {
		dNumbers[index] = transformer(number)
	}
	return dNumbers
}

func getTransFormer(numbers *[]int) transformerFunction {
	if (*numbers)[0] == 1 {
		return double
	}
	return triple
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
