package main

import "fmt"

/* Struct adalah tipe data seperti tipe data lainnya
dia bisa digunakan sebagai parameter untuk function
Namun jika kita ingin menambahkan method ke dalam structs,
sehingga seakan" struct memiliki function
Method adalah function
*/

type Customer struct {
	Name, address string
	Age           int
	isMarried     bool
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func (cus Customer) sayHu() {
	fmt.Println("huuu from", cus.Name)
}

func main() {
	var customer Customer
	customer.Name = "Muhammad febri andani"
	customer.address = "Jalan cirendeu raya"
	customer.Age = 22
	customer.isMarried = false

	customer.sayHi("Joko")
	customer.sayHu()

	// fmt.Println(customer)

	// //program struct literals
	// customer2 := Customer{
	// 	Name:      "Andi gunawan",
	// 	address:   "Indonesia",
	// 	Age:       23,
	// 	isMarried: false,
	// }

	// fmt.Println("Name : ", customer2.Name)
	// fmt.Println("Address : ", customer2.address)
	// fmt.Println("Age : ", customer2.Age)
	// if customer2.isMarried {
	// 	fmt.Println("isMarried : Menikah")
	// } else {
	// 	fmt.Println("isMarried : Belum menikah")
	// }

}
