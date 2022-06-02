package main

import "fmt"

func main() {
	type NoKtp string
	type Married bool

	var ktpFebri NoKtp = "212412414"
	fmt.Println(ktpFebri)
	fmt.Println(NoKtp("141241948151"))
	fmt.Println("==============")
	var marriedStatus Married = true
	fmt.Println(marriedStatus)
}
