package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktiber",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	fmt.Println(append(slice1, "febri"))

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "febri")
	fmt.Println(slice3)
	slice3[1] = "Bukan desember"
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(months)
}
