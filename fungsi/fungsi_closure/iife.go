package main

import "fmt"

//program Immediately - invoked function expression(IIFE)
func main(){
	var numbers = []int{2,4,5,2,0,2,10,6,4,2}

	var newNum = func(min int)[]int{
		var r []int
		for _, e := range numbers{
			if e < min{
				continue
			}
			r = append(r,e)
		}
		return r
	}(3)

	fmt.Println("original number :",numbers)
	fmt.Println("filtered number :",newNum)
}