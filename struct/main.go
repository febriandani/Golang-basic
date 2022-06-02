package main

import "fmt"

/* Struct adalah template data yang digunakan untuk
menggabungkan nol atau lebih tipe data lainnya dalam satu kesatuan
Struct biasanya representasi data dalam program aplikasi yang kita buat
Data di struct disimpan dalam field
Sederhananya struct adalah kumpulan dari field
*/

/* struct adalah template data atau prototype data
struct tidak bisa langsung digunakan
namun bisa membuat data/object dari struct yang telah kita buat
*/

type Customer struct {
	Name, address string
	Age           int
	isMarried     bool
}

func main() {
	var customer Customer
	customer.Name = "Muhammad febri andani"
	customer.address = "Jalan cirendeu raya"
	customer.Age = 22
	customer.isMarried = false

	fmt.Println(customer)

	//program struct literals
	customer2 := Customer{
		Name:      "Andi gunawan",
		address:   "Indonesia",
		Age:       23,
		isMarried: false,
	}

	fmt.Println("Name : ", customer2.Name)
	fmt.Println("Address : ", customer2.address)
	fmt.Println("Age : ", customer2.Age)
	if customer2.isMarried {
		fmt.Println("isMarried : Menikah")
	} else {
		fmt.Println("isMarried : Belum menikah")
	}

}
