package main

import (
	"fmt"
)

func main() {
	name := "lia"

	switch name {
	case "febri":
		fmt.Println("Hello febri")
	case "Joko":
		fmt.Println("Hello joko")
	default:
		fmt.Printf("Hi %s, Salam kenal", name)
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("\nNama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	length2 := len(name)
	switch {
	case length2 > 10:
		fmt.Println("Nama kepanjangan")
	case length2 > 5:
		fmt.Println("Nama sudah cukup")
	default:
		fmt.Println("Nama kependekan")
	}
}
