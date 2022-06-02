package main

import "fmt"

type student struct{
	nama string
	age int
}

func main(){
	var s1 student
	var nama string
	var age int

	fmt.Println("Nama : ",s1.nama)
	fmt.Println("Age : ",s1.age)

	var s2 = student{"ethan",21}

	var s3 = student{"jaka",23}

	fmt.Printf("Nama : %s\nAge : %d",s2.nama,s2.age)
	fmt.Printf("\nNama : %s\nAge : %d",s3.nama,s3.age)
}