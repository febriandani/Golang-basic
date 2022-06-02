package main

import "fmt"

func main(){
	var positivenum uint8 = 89
	var negativenum = -12345644

	fmt.Println("Tipe data int / numerik non desimal")
	fmt.Println("Bilangan positif = ",positivenum)
	fmt.Println("Bilangan negatif = ",negativenum)

	fmt.Println("\nTipe data desimal")
	var decimalnum = 4.5

	fmt.Printf("Bilangan desimal = %f",decimalnum)
	fmt.Printf("\nBilangan desimal = %.3f",decimalnum)

	fmt.Println("Tipe data boolean")
	var hasil bool = true
	
	fmt.Printf("\nhasil? %t \n", hasil)

	fmt.Println("\nTipe data string")
	var message string = "halo"
	fmt.Println("\nIsi string = ",message)

	var message2 string = `nama saya"Muhammad febri andani".
	Salam kenal.
	mari belajar "golang".`

	fmt.Println(message2)

}