package main

import (
	"fmt"
	"strings"
)

func main() {
	avg := calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	msg := fmt.Sprintf("Rata - rata : %.2f",avg)

	fmt.Println(msg)

	fmt.Println("\n\nYourhobbiesFunc")
	//bagian ini menggunakan parameter biasa
	// yourhobbies("wick","sleeping","badminton")
	//sedangkan ini menggunakan data slice dan menggunakan tanda titik 3 kali
	var hobbies = []string{"sleeping","badminton"}
	yourhobbies("wick", hobbies...)
}


//func calculate variadic
func calculate(numbers ...int) float64 {
	total := 0
	for _, number := range numbers {
		total += number
	}

	avg := float64(total) / float64(len(numbers))
	return avg

}

//func yourhobbies dengan parameter biasa dan variadic
func yourhobbies (nama string, hobby ...string){
	var hoobiesAsString = strings.Join(hobby,",")

	fmt.Printf("hello, my name is :%s\n",nama)
	fmt.Printf("my hobbies are :%s\n",hoobiesAsString)
}