package main

import "fmt"

const NMax int = 1000

type tabInt [NMax]int

func main() {
	var N int
	var arr tabInt
	fmt.Scan(&N)
	isiArray(&arr, N)
	prosesAngkaGanjil(arr, N)
}

func isiArray(arrInt *tabInt, n int) {
	/*I.S. Terdefinisi nilai bilangan bulat sebanyak n.
	  F.S. Array arrInt berisi data yang diberikan*/
	var x, i int //variabel untuk input dan variabel untuk loop
	for i = 0; i < n; i++ {
		fmt.Scan(&x)
		for x < 0 { // loop untuk mengecek apakah inputan < 0 , jika benar maka meminta inputan ulang
			fmt.Scan(&x)
		}
		arrInt[i] = x
	}
}

func cekGanjil(x int) bool {
	/*Mengembalikan true apabila x adalah bilangan ganjil dan false jika sebaliknya*/
	return x%2 != 0
}

func prosesAngkaGanjil(arrInt tabInt, n int) {
	/*I.S. Terdefinisi array arrInt, array tidak kosong.
	  F.S. Menampilkan semua bilangan ganjil dalam array dan bilangan ganjil terbesar dalam array arrInt*/
	var i, maxGanjil int //variabel untuk perulangan dan untuk menyimpan nilai tertinggi dari kumpulan angka ganjil
	maxGanjil = 0
	for i = 0; i < n; i++ {
		if cekGanjil(arrInt[i]) {
			fmt.Print(arrInt[i], " ")
			if maxGanjil < arrInt[i] {
				maxGanjil = arrInt[i]
			}
		}
	}
	if maxGanjil == 0 {
		fmt.Print("-")
	}
	fmt.Println()
	fmt.Print(maxGanjil)
}
