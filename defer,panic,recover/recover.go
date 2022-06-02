package main

import "fmt"

/* Recover adalah function yang bisa kita gunakan untuk menangkap data panic
dengan revocer proses panic terhenti, sehingga program akan
tetap berjalan */

/* Menurut pemikiran saya
jadi ketika menggunakan recover string yang di dalam panic tidak akan di tampilkan
akan tetapi ditampilkan menggunakan func revocer */

/*jadi ketika ada error panic, panic tersebut telah ditangani oleh recover
maka error telah di tangani program nya akan tetap berjalan */

func endApp() {
	messageErr := recover()
	if messageErr != nil {
		fmt.Println("Error dengan message :", messageErr)
	}
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi error")
	}

	fmt.Println("Aplikasi berjalan lancar")
}

func main() {
	runApp(false)
}
