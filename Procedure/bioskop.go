package main

import "fmt"

type film struct {
	nama_film, nama_hari string
	tanggal_merah        bool
}

func main() {
	var tiket film

	fmt.Scan(&tiket.nama_film, &tiket.nama_hari, &tiket.tanggal_merah)
	hitungHarga(tiket)
}
func hitungHarga(tiket film) {

	if tiket.tanggal_merah == true || tiket.nama_hari == "Sabtu" || tiket.nama_hari == "Minggu" {
		hargaAkhir := 50000 + (50000 * 0.50)
		fmt.Println(hargaAkhir)
	} else {
		hargaAkhir := 50000
		fmt.Println(hargaAkhir)
	}
}
