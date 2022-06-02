package main

import (
	"fmt"
)

func main()  {
	const firstname string = "Dani"
	const lastname = "andin"

	fmt.Println(firstname)
	fmt.Println(lastname)


	//multiple constant
	const(
		country = "indonesia"
		name = "salva"
		age = 22
	)

	fmt.Println(country)
	fmt.Println(name)
	fmt.Println(age)
}