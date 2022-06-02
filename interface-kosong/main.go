package main

import "fmt"

/*
interface kosong adalah interface yang tidak memiliki
deklarasi method satupun, hal ini membuat secar otomatis
semua tipe data akan menjadi implementasinya
*/

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func main() {
	kosong := Ups(1)

	fmt.Println(kosong)
}
