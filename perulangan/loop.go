package main

import "fmt"

func main() {
	fmt.Println("Perulangan menggunakan key for")

	for i := 0; i <= 10; i++ {
		fmt.Println("angka ",i)
	}

	fmt.Println("Atau bisa dengan")

	var a = 1

	for a <= 5 {
		fmt.Println("Angka", a)
		a++
	}

	fmt.Println("Atau bisa menggunakan break & continue")

	for i := 0; i < 10; i++ {
		if i % 2 == 1{
			continue
		} 

		if i > 8 {
			break
		}

		fmt.Println("Angka ",i  )
	}
}