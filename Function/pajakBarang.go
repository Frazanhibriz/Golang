package main

import "fmt"

func hargaBarangkenaPajak(asal string, hbarang int) float64 {
	hbarang2 := float64(hbarang)

	if asal == "dalam" {
		hbarang2 = hbarang2 * (100 - 8) / 100
	} else if asal == "luar" {
		hbarang2 = hbarang2 * (100 - 18) / 100
	}

	return hbarang2
}

func main() {
	var asal string
	var hbarang int

	fmt.Scan(&asal, &hbarang)
	fmt.Printf("%.2f", hargaBarangkenaPajak(asal, hbarang))
}
