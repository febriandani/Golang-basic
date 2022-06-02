package main

import "fmt"

//function dengan parameter
func sayHello(name string, lastname string, usia int) {
	fmt.Println("hello", name, lastname, usia)
}

func main() {
	sayHello("febri", "andani", 22)
}
