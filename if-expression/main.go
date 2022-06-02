package main

import "fmt"

func main() {
	var nilai float32

	fmt.Println("Masukan nilai : ")
	fmt.Scan(&nilai)

	if nilai > 90 {
		fmt.Println("LULUS GRADE A")
	} else if nilai > 80 {
		fmt.Println("LULUS GRADE B")
	} else if nilai > 60 {
		fmt.Println("LULUS GRADE C")
	} else if nilai > 40 {
		fmt.Println("LULUS GRADE D")
	} else {
		fmt.Println("TIDAK LULUS, GRADE E")
	}
}
