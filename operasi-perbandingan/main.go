package main

import "fmt"

func main() {
	var name1 = "Eko"
	var name2 = "Eko"

	// bool
	var result = name1 == name2
	fmt.Println(result)

	var val1 = 100
	var val2 = 200

	fmt.Println(val1 > val2)
	fmt.Println(val1 < val2)
	fmt.Println(val1 == val2)
	fmt.Println(val1 != val2)
}
