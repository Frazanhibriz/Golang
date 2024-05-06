package main

import "fmt"

func main() {
	var n, angka int
	fmt.Scan(&n, &angka)
	printPolaSegitiga(1, 1, n, angka)
}

func printPolaSegitiga(baris int, kolom int, n int, angka int) {
	/*  I.S Terdefinisi nilai bilangan bulat baris, kolom, dan n
	F.S menampilkan pola  string "*" yang berbentuk segitiga */
	if kolom <= n {
		if kolom < n-(baris-1) { // Pengecekan Pengulangan kolom dan baris dengan operasi perbandingan,
			fmt.Print(" ")
		} else {
			fmt.Print(angka)
			angka = angka + 1
		}
		printPolaSegitiga(baris, kolom+1, n, angka%10) // gunakan fungsi rekursif pada baris ini  hint : gunakan modulus 10

	} else if baris < n {
		fmt.Println()

		printPolaSegitiga(baris+1, 1, n, angka) // gunakan fungsi rekursif pada baris ini
	}
}
