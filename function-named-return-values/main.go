package main

import "fmt"

func getCompleteName() (fname, midname, lname string) {
	fname = "muhammad"
	midname = "febri"
	lname = "andani"

	return
}

func main() {
	_, midname, lname := getCompleteName()
	fmt.Println(midname, lname)

}
