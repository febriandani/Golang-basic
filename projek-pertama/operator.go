package main

import "fmt"

func main()  {
	fmt.Println("Operator aritmatika")
	a := 36
	b := 5

	var value = ((a + b) * 4 - 2)
	var isEqual = (value == 5)
	fmt.Println(value)
	fmt.Printf("Nilai dari value (5) = %d (%t) \n",value,isEqual)
	fmt.Printf("nilai %d (%t) \n", value, isEqual)
}