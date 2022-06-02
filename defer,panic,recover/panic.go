package main

import "fmt"

/* Panic func adalah func yang bisa kita gunakan
utnuk menghentikan program
panic func biasanya di panggil ketika terjadi error
pada saat program kita berjalan
saat panic func dipanggil, program akan terhenti
namun defer func tetap akan dieksekusi */

func endApp() {
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
	runApp(true)
}
