package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	printPolaBilangan(1, 1, n)
}

func printPolaBilangan(baris, kolom, n int) {
	/*  I.S bilangan bulat baris, kolom, dan n
	F.S menampilkan pola bilangan bulat yang menyatakan baris hingga membentuk pola X yang berukuran N × N */
	if n >= baris {
		if n >= kolom {
			if baris == kolom || kolom == (n-baris)+1 { // Pengecekan Pengulangan kolom dan baris dengan operasi perbandingan,
				// hint : terdapat beberapa operasi perbandingan, yaiut == dan ||
				fmt.Print(baris)
			} else {
				fmt.Print(" ")
			}
			printPolaBilangan(baris, kolom+1, n) // gunakan fungsi rekursif pada baris ini
		} else {
			fmt.Println()
			printPolaBilangan(baris+1, 1, n) // gunakan fungsi rekursif pada baris ini
		}
	}
}
