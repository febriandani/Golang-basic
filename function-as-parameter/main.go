package main

import "fmt"

func sayHelloWithFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "Kata tidak pantas!!!"
	} else {
		return name
	}
}

// atau cara lain function type declaration
type Filter2 func(string) string

func sayHelloWithFilter2(name string, filter2 Filter2) {
	fmt.Println("Hello", filter2(name))
}

func spamFilter2(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Febri", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)
	sayHelloWithFilter2("Anjing", spamFilter2)

}
