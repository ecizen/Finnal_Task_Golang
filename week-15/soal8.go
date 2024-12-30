package main

import "fmt"

func main() {
	var jumlahRenang , Jatahminggu int

	fmt.Print("Masukkan jumlah renang: ")
	fmt.Scan(&jumlahRenang)

	Jatahminggu = jumlahRenang / 7

	if jumlahRenang % 7 == 0 {
		fmt.Println("Jatah Renang Yuki habis di minggu ke-",Jatahminggu)
	} else {
		fmt.Println("Jatah Renang Yuki habis di minggu ke-",Jatahminggu + 1)
	}

}