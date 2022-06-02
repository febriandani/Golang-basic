// closure adalah kemampuan sebuah function berinteraksi dengan data
// disekitarnya dalam scope yang sama

package main

import "fmt"

func main() {
	name := "dani"
	counter := 0
	increment := func() {
		name = "febri"
		fmt.Println("increment")
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)
}
