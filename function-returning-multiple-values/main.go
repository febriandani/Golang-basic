package main

import "fmt"

func getFullName() (string, string, int) {
	return "Muhammad", "febri", 12
}

func main() {
	fmt.Println(getFullName())
	fname, lname, usia := getFullName()
	fmt.Println("Nama depan : ", fname)
	fmt.Println("Nama belakang : ", lname)
	fmt.Println("Usia : ", usia)

	_, _, usia2 := getFullName()
	fmt.Println("Usia : ", usia2)
}
