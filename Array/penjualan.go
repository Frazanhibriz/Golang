package main

import (
	"fmt"
)

type array12Int [12]int

func inputPenjualan(arrPenjualan *array12Int) {
	/*I.S. Terdefinisi arrPenjualan
	F.S. arrPenjualan berisi dengan data penjualan tiap bulan dari tahun sebelumnya*/
	fmt.Print("Input data penjualan (12 bulan): ")
	for i := 0; i < 12; i++ {
		fmt.Scan(&arrPenjualan[i])
	}
}

func avgPenjualan(arrPenjualan array12Int) float64 {
	//mengembalikan rata-rata dari data penjualan
	var total int
	for i := 0; i < 12; i++ {
		total += arrPenjualan[i]
	}
	return float64(total) / 12
}

func bulanTerbanyak(arrPenjualan array12Int) int {
	//mengembalikan bulan dengan total penjualan tertinggi
	var maxx, bulanmaxx int
	for i := 0; i < 12; i++ {
		if arrPenjualan[i] > maxx {
			maxx = arrPenjualan[i]
			bulanmaxx = i + 1
		}
	}
	return bulanmaxx
}

func bulanTersedikit(arrPenjualan array12Int) int {
	//mengembalikan bulan dengan total penjualan terendah
	var minn, bulanminn int
	minn = arrPenjualan[0]
	bulanminn = 1
	for i := 1; i < 12; i++ {
		if arrPenjualan[i] < minn {
			minn = arrPenjualan[i]
			bulanminn = i + 1 
		}
	}
	return bulanminn
}

func main() {
	var arrPenjualan array12Int
	inputPenjualan(&arrPenjualan)
	fmt.Println("Rata-rata penjualan tiap bulan: ", avgPenjualan(arrPenjualan))
	fmt.Println("Bulan dengan penjualan tertinggi: ", bulanTerbanyak(arrPenjualan))
	fmt.Println("Bulan dengan penjualan terendah: ", bulanTersedikit(arrPenjualan))
}
