package main

import "fmt"

func jumlahKelipatan(bil1 int, bil2 int, x int) int {
	var jumlah int

	for i := bil1; i <= bil2; i++ {
		if i%x == 0 {
			jumlah = jumlah + i
		}
	}

	return jumlah
}

func main() {
	var bil1, bil2, x int

	fmt.Scan(&bil1, &bil2, &x)
	fmt.Println(jumlahKelipatan(bil1, bil2, x))
}
