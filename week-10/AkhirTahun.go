package main

import "fmt"

func main() {
	var totalBelanja, disc, cashBack, hargaTotal float64
	var member bool

	// Input
	fmt.Print("Masukkan total belanja: ")
	fmt.Scan(&totalBelanja)
	fmt.Print("Apakah member (true/false): ")
	fmt.Scan(&member)


	disc = 0.1
	cashBack = 75000


	fmt.Println("Kartu?", member)
	fmt.Println("Diskon?", totalBelanja >= 100000)
	fmt.Println("Cashback?", member && totalBelanja >= 200000)

	
	if totalBelanja >= 100000 {
		hargaTotal = totalBelanja - (totalBelanja * disc) 

		if member && totalBelanja >= 200000 {
			hargaTotal -= cashBack 
		}
	} else {
		hargaTotal = totalBelanja
	}

	// Output hasil
	fmt.Printf("Harga total setelah diskon dan cashback: %.0f\n", hargaTotal)
}
