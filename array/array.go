package main

import "fmt"

func main() {
	//cara horizontal
	var buah = [4]string{"Apel", "mangga", "nanas", "Jambu"}

	//cara vertikal dan tanpa jumlah [elemen] / [...]
	var uang = [...]int64{
		1000,
		2000,
		5000,
		10000,
	}

	fmt.Println("Jumlah len \t",len(buah))
	fmt.Println("Isi \t",buah)

	fmt.Println("Jumlah len \t ",len(uang))
	fmt.Println("Isi Uang \t",uang)

	// //menggunakan keywords for
	// fmt.Println("Menggunakan format for")
	// for i := 0; i <= len(buah); i++ {
	// 	fmt.Printf("elemen ke %d : %s\n", i, buah[i])
	// }

	fmt.Println("Menggunakan format for - range")
	for i, buah := range buah{
		fmt.Printf("Elemen ke %d : %s\n",i,buah)
	}
}