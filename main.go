package main

import "fmt"

func main() {
	//numbers := []int{1, 2, 3, 4, 5}
	sum := sumup(5, 1, 10, 15)
	fmt.Println(sum)

	sum2 := sumup(7, 1, 2, 3, 4, 5)
	fmt.Println(sum2)
}

func sumup(firstValue int, numbers ...int) int {
	fmt.Println(firstValue)
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}
