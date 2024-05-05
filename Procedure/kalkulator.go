package main

import "fmt"

func main() {
	var pilih int

	for {
		menu()
		fmt.Print("pilih (1/2/3/4)?")
		fmt.Scan(&pilih)
		if pilih == 1 {
			hitungJumlah()
		} else if pilih == 2 {
			hitungKali()
		} else if pilih == 3 {
			hitungBagi()
		} else if pilih == 4 {
			break
		}
	}
}

func menu() {
	fmt.Println("-----------------------")
	fmt.Println("         MENU          ")
	fmt.Println("-----------------------")
	fmt.Println("1. Hitung Penjumlahan")
	fmt.Println("2. Hitung Perkalian")
	fmt.Println("3. Hitung Pembagian")
	fmt.Println("4. Exit")
	fmt.Println("-----------------------")
}

func hitungJumlah() {
	var n, m, hasilJumlah int

	fmt.Print("Masukkan dua bilangan yang akan dijumlahkan:")
	fmt.Scan(&n, &m)
	hasilJumlah = n + m
	fmt.Print("Hasil penjumlahan:")
	fmt.Println(hasilJumlah)
}

func hitungKali() {
	var n, m, hasilKali int

	fmt.Print("Masukkan dua bilangan yang akan dikalikan:")
	fmt.Scan(&n, &m)
	hasilKali = n * m
	fmt.Print("Hasil perkalian:")
	fmt.Println(hasilKali)
}

func hitungBagi() {
	var n, m, hasilBagi float64

	fmt.Print("Masukkan dua bilangan yang akan dibagi:")
	fmt.Scan(&n, &m)
	hasilBagi = n / m
	fmt.Print("Hasil pembagian:")
	fmt.Println(hasilBagi)
}
