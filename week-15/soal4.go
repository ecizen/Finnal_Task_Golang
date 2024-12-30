package main

import "fmt"

func main() {
	var x, y, z float64
	var belanja, diskon float64

	fmt.Scan(&x, &y, &z)

	diskon = z * x /100 

	if diskon > y {
		diskon = y
	}
	belanja = z - diskon
	fmt.Println(belanja)

}