package main

import "fmt"

func main() {
	var a = 10
	var b = 20
	var c = a + b
	fmt.Println(c)

	var result = 10 * 2
	fmt.Println(result)

	var mod float32 = 101 % 3
	fmt.Printf("%f", mod)

	fmt.Println("\n====UNARY=======")
	var i = 10
	i += 10
	fmt.Println(i)

	i++
	fmt.Println(i)

	var negative = -10
	var positive = 100
	fmt.Println(negative)
	fmt.Println(positive)
}
