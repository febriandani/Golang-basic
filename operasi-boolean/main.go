package main

import "fmt"

func main() {
	var nilaiAkhir = 80.1
	var absensi = 80.1

	var lulusNilaiAkhir = nilaiAkhir > 80
	var lulusAbsensi = absensi > 80

	fmt.Println(lulusNilaiAkhir)
	fmt.Println(lulusAbsensi)

	var lulus = lulusNilaiAkhir && lulusAbsensi
	fmt.Println(lulus)
}
