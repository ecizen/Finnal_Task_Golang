package main

import "fmt"

func main() {
	var bulanIni, bulanLalu, profit float64

	fmt.Print("Masukkan profit bulan lalu:")
	fmt.Scan(&bulanLalu)
	fmt.Print("Masukkan profit bulan ini:")
	fmt.Scan(&bulanIni)

	if bulanIni > bulanLalu {
		profit = bulanIni - bulanLalu
		fmt.Println("Peningkatan profit sebanyak", profit)
	} else if bulanLalu > bulanIni {
		profit = bulanLalu - bulanIni
		fmt.Println("Penurunan profit sebanyak", profit)
	} else {
		fmt.Println("Tidak ada profit")
	}
}