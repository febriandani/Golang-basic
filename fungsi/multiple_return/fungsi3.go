package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Multiple return")
	var d float64 = 15

	var area, kll = hitung(d)
	fmt.Printf("\nLuas lingkaran : %.2f",area)
	fmt.Printf("\nKeliling lingkaran : %.2f",kll)

	fmt.Println("\n\nFungsi dengan Predefined return value")
	
	var area2, kll2 = calculate(d)
	fmt.Printf("\nLuas lingkaran : %.2f",area2)
	fmt.Printf("\nKeliling lingkaran : %.2f",kll2)
}



//func multiple return
func hitung(d float64)(float64,float64){
	//hitung luas
	area := math.Pi * math.Pow(d / 2,2)
	//hitung keliling
	kll := math.Pi * d

	return area, kll
}

//func Predefined return value
func calculate(d float64) (area float64, kll float64){
	area = math.Pi * math.Pow(d/2,2)
	kll = math.Pi * d	

	return
}