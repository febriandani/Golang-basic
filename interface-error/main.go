package main

import (
	"errors"
	"fmt"
)

func pembagi(nilai float32, pembagi float32) (float32, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi must be > 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {
	hasil, err := pembagi(100, 0)
	if err != nil {
		fmt.Println("Error ", err)
	} else {
		fmt.Printf("Hasil: %v\n", hasil)
	}
}
