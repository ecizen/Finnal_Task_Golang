package main

import "fmt"

func main() {
	var x1, x2 int

	fmt.Scan(&x1)
	fmt.Scan(&x2)

	status := x1 % 2 == 0 && x2 % 2 != 0 || x1 % 2 !=0 && x2 % 2 == 0  
	if status {
		println("berhasil")
	}

}