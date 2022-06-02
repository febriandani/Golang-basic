package main

import "fmt"

func main() {
	var nilai32 int32 = 1299999999
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println("==========PROGRAM KONVERSI TIPE DATA INT 32 TO INT 64,16,8============")
	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)
	fmt.Println(nilai8)

	fmt.Println("==========PROGRAM KONVERSI TIPE STRING TO BYTE and BYTE TO ============")
	var name = "Muhammad Febri Andani"
	var e = name[0]
	var eString = string(e)

	fmt.Println("string to byte")
	fmt.Println(e)

	fmt.Println("byte to string")
	fmt.Println(eString)
}
