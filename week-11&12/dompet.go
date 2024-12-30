package main

import "fmt"

func main() {
	var saldo, transaksi int


	fmt.Scan(&transaksi)
	saldo = transaksi

	// Loop untuk memproses transaksi berikutnya
	for transaksi != 0 {
		// Membaca transaksi berikutnya
		fmt.Scan(&transaksi)

		// Jika transaksi bukan 0, tambahkan atau kurangi saldo sesuai dengan transaksi
		saldo += transaksi
	}

	// Menampilkan saldo akhir
	fmt.Println(saldo)
}
