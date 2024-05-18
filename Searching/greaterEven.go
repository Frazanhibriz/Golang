package main

import "fmt"

const N int = 1000

func main() {
	var T [N]int
	var jumlah int
	fmt.Scanln(&jumlah)
	for i := 0; i < jumlah; i++ {
		fmt.Scan(&T[i])
	}
	fmt.Println(genapTerbesar(T, jumlah))
}

func genapTerbesar(T [N]int, jumlah int) int {
	//mengembalikan bilangan genap terbesar dalam array T dengan banyak data sebanyak jumlah
	var terbesar int

	terbesar = -1

	for i := 0; i < jumlah; i++ {
		if T[i]%2 == 0 && T[i] > terbesar {
			terbesar = T[i]
		}
	}
	return terbesar

}
