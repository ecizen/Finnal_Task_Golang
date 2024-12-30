package main

import "fmt"

func main() {
	var angka int
	fmt.Scan(&angka)

	if angka%3 == 0 {
		fmt.Println("Kelipatan 3")
	}

	if angka%5 == 0 {
		fmt.Println("Kelipatan 5")
	}
}