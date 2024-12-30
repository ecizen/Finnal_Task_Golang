package main

import "fmt"

func main() {
	var (
		tarif, potongan, diskon, tarif_awal float64
		durasi, kelebihan int
		member string
	)

	
	fmt.Scan(&member, &durasi)

	if member == "Gold" {
		diskon = 0.5
	} else if member == "Silver" {
		diskon = 0.25
	} else {
		diskon = 0
	}

	
	if durasi <= 2 {
		tarif_awal = 65000 * float64(durasi)
	} else {
		kelebihan = durasi - 2
		tarif_awal = 65000*2 + 20000*float64(kelebihan)
	}

	potongan = diskon * tarif_awal
	tarif = tarif_awal - potongan
	fmt.Printf("IDR %.0f\n", tarif)
}
