package main

import "fmt"

func main() {
	var n int
	var maxDigit int

	fmt.Println("Masukkan bilangan bulat positif:")
	fmt.Scan(&n)

	
	for n > 0 {
		digit := n % 10 
		if digit > maxDigit {
			maxDigit = digit 
		}
		n /= 10 
	}

	fmt.Println(maxDigit)
}
