package main

import "fmt"

func diskon(totalBelanja int) int {

	/*Mengembalikan angka 1 jika belanja
	pembeli memenuhi syarat diskon dan
	angka 0 jika tidak memenuhi syarat*/
	n := 4
	if totalBelanja%100 == 0 {
		return 1
	} else if totalBelanja%n == 0 {
		return 1
	} else {
		return 0
	}
}

func cashback(totalBelanja int) int {

	/*Mengembalikan angka 1 jika belanja pembeli
	memenuhi syarat cashback dan
	angka 0 jika tidak memenuhi syarat*/
	n := 3
	if totalBelanja%100 == 0 {
		return 1
	} else if totalBelanja%n == 0 {
		return 1
	} else {
		return 0
	}
}

func main() {
	var banyaknyaPembeli, totalBelanja, i int
	var DapatDiskon, DapatCashback int
	fmt.Scan(&banyaknyaPembeli)
	DapatDiskon = 0
	DapatCashback = 0
	for i = 1; i <= banyaknyaPembeli; i++ {
		fmt.Scan(&totalBelanja)
		DapatDiskon = DapatDiskon + diskon(totalBelanja)
		DapatCashback = DapatCashback + cashback(totalBelanja)
	}
	fmt.Println(DapatDiskon, DapatCashback)
}
