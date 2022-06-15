package main

import "fmt"

/* secara default di golang semua variable itu dipassing by value
bukan by reference, artinya jika kita mengirim sebuah variable ke dalam function,
method atau variable lain, sebenarnya dikirim adalah duplikasi valuenya */

/* dengan menggunakan pointer kemampuan membuat reference
ke lokasi data di memory yang sama, tanpa menduplikasi data yang sudah ada
Sederhananya debgan kemampuan pointer, kita bisa membuat pass by reference*/

/*Kalau pass by value datanya di duplikat, sedangkan by reference tidak*/
/*Untuk membuat sebuah variable dengan nilai pointer ke variable lain, kita bisa menggunakan
operator & diikuti dengan nam var nya*/
type Address struct {
	City     string
	Province string
	Country  string
}

func main() {
	address1 := Address{}
	address1.City = "Subang"
	address1.Province = "Jawa Barat"
	address1.Country = "Indonesia"

	address2 := &address1
	// address3 := &address1
	address2.City = "Bandung"
	//mengubah isi dari pointer di awal
	address2 = &Address{"Jakarta", "DKI JAKARTA", "Indonesia"}
	// buat pointer baru
	// *address3 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println("Address 1 adalah value awal diubah citynya : ", address1)
	fmt.Println("Address 2 adalah mengubah isi dari pointer di awal : ", address2)
	// fmt.Println("Address 3 Mengubah seluruh data di pointer awal :", address3)

	address4 := new(Address)
	address4.City = "Tangerang"
	fmt.Println(address4)
}
