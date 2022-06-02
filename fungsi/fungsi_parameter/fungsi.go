package main

import (
	"fmt"
	"strings"
)

//penerapan fungsi sebagai parameter
func filter(data []string, callback func(string) bool) []string{
	var result []string
	for _, each := range data{
		if filtered := callback(each); filtered{
			result = append(result, each)
		}
	}
	
	return result
}

func main(){
	var data = []string{"wick","jason","ethan"}
	var dataContainsO = filter(data, func(each string) bool {
		return strings.Contains(each,"o")
	})
	var dataLen5 = filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("data aslinya : ",data)
	fmt.Println("data mencari huruf o : ",dataContainsO)
	fmt.Println("data filter jumlah huruf \"5\"\t:",dataLen5)
}