package main

import "fmt"

func main() {
	count := 1

	fmt.Println("for")
	for count <= 10 {
		fmt.Println("*", count)
		count++
	}

	fmt.Println("For dengan statement")
	for count2 := 1; count2 <= 10; count2++ {
		fmt.Println("*", count2)
	}

	fmt.Println("For dengan range untuk melakukan iterasi\nterhadap semua data collection\ncontohnya array, slice, map")
	fmt.Println("COBA DI SLICE")
	slice := []string{"muhammad", "febri", "andani"}

	// for i, user := range slice {
	// 	fmt.Println("index", i, "name", user)
	// }
	for _, user := range slice {
		fmt.Println(user)
	}

	fmt.Println("COBA DI MAP")
	person := make(map[string]string)
	person["name"] = "Febri"
	person["job"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
