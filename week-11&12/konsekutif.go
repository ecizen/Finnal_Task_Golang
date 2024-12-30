package main

import "fmt"

func main() {
	var X int
	fmt.Scan(&X)

	isKonsekutif := true
	DigitTerakhir := X % 10
	X = X / 10

	for X > 0 {
		DigitSekarang := X % 10
		if DigitTerakhir-DigitSekarang != 1 &&DigitSekarang-DigitTerakhir != 1 {
			isKonsekutif = false
			break
		}
		DigitTerakhir = DigitSekarang
		X = X / 10
	}

	if isKonsekutif {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
