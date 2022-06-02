package main

import "fmt"

/*defer function adalah func yang bisa kita jadwalkan untuk dieksekusi
setelah sebuah func selesai di eksekusi
defer func akan selalu di eksekusi walaupun terjadi error di func yang di eksekusi
*/
//tujuannya ketika ada func yang ingin dieksekusi ketika terjadi error

func logging() {
	fmt.Println("selesai memanggil function")
}

func runApp(val int) {
	defer logging()
	fmt.Println("Run Application")
	result := 10 / val
	fmt.Println(result)
	// logging()

}

func main() {
	runApp(5)
}
