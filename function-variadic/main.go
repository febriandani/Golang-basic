package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	total := sumAll(10, 15, 20, 25, 30, 5)
	fmt.Println(total)

	numberSlice := []int{20, 10, 20, 30, 40, 20}
	total2 := sumAll(numberSlice...)
	fmt.Println(total2)
}
