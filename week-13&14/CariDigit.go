package main

import "fmt"

func main() {
	var x, n int
	var hasil bool

	
	fmt.Println("Masukkan dua bilangan bulat positif x :")
	fmt.Scan(&x, &n)

	for n > 0 {
		digit := n % 10 
		if digit == x {
			hasil = true
			break
		}
		n /= 10 
	}

	// Output hasil
	if hasil {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
