package main

import "fmt"

func main() {
	person := map[string]string{
		"name":   "febri",
		"alamat": "poncol",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["alamat"])

	var book = make(map[string]string)
	book["title"] = "Belajar golang"
	book["author"] = "eko"
	book["ups"] = "salah"
	fmt.Println(book)

	delete(book, "ups")
	fmt.Println(book)
}
