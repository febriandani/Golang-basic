package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	var nama = []string{"john","wick"}
	printMsg("halo",nama)


	rand.Seed(time.Now().Unix())
	var randomVal int

	randomVal = randomWithRange(676,898)
	fmt.Println("Random number : ",randomVal)
}

func randomWithRange(min, max int) int{
	var val = rand.Int() % (max-min + 1) +min
	return val
}

func printMsg(message, arr[]string){
	var namaString = strings.Join(arr,"")
	fmt.Println(message, namaString)
}