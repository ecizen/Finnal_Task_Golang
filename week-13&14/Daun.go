package main

import "fmt"

func main() {
	var n, lebar, maxLebar int

	fmt.Println("Masukkan jumlah daun dan lebar:")
	fmt.Scan(&n)


	for i := 0; i < n; i++ {
		fmt.Scan(&lebar)
		if lebar > maxLebar {
			maxLebar = lebar
		}
	}

	fmt.Println(maxLebar)
}
