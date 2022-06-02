package main

import "fmt"

func main() {
	divideNum(10, 2)
	divideNum(8, 0)
	divideNum(100, -2)

}

func divideNum(m, n int) {
	if n == 0 {
		fmt.Printf("invalide divider, %d cannot divided by %d\n",m,n)
		return
	}

	res := m / n
	fmt.Printf("%d / %d = %d \n",m,n,res )
}