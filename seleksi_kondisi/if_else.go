package main

import "fmt"

func main(){
	var point = 3

	if point== 10{
		fmt.Println("Lulus dengan sempurna")
	} else if point > 5{
		fmt.Println("Lulus")
	} else if point > 4{
		fmt.Println("Hampir lulus")
	} else{
		fmt.Printf("Tidak lulus. Nilai anda adalah : %d ",point)
	}
}