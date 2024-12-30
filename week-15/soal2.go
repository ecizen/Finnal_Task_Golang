package main

import "fmt"

func main() {
	var jamDatang, menitDatang, lamaService int

	fmt.Scan(&jamDatang, &menitDatang, &lamaService)

	jamTutup := 20 * 60 

	durasi := jamDatang  * 60 + menitDatang + lamaService

	jamSelesai := durasi / 60
	menitSelesai := durasi % 60

	if durasi < jamTutup || jamSelesai == 20   && menitSelesai == 0 {
		fmt.Printf("%02d %02d\n", jamSelesai, menitSelesai)
	} else {
		fmt.Println("Silahkan datang kembali besok")
	}

	
}
