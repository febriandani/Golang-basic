package main

import (
	"fmt"
)

func random() interface{} {
	return 12345644
}

func main() {
	result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	case float32:
		fmt.Println("Float", value)
	case bool:
		fmt.Println("Bool", value)
	default:
		fmt.Println("Invalid", value)
	}
}
