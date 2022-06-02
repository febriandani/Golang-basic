package main

import "fmt"

func main() {
	var name [3]string
	name[0] = "muhammad"
	name[1] = "febri"
	name[2] = "andani"

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	var umur = [3]int{
		12, 13, 15,
	}

	fmt.Println(umur[0])
	fmt.Println(umur[1])
	fmt.Println(umur[2])
	fmt.Println(umur)
	fmt.Println(len(umur))
	umur[0] = 19
	fmt.Println(umur[0])
}
