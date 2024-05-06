package main

import "fmt"

const NMAX int = 10

type tabInt [NMAX]int

func main() {
	var data tabInt
	var nData int
	baca(&data, &nData)
	cetak(data, nData)
	fmt.Println(jumlah(data, nData), rata_rata(data, nData))
}

func baca(A *tabInt, n *int) {
	var i int
	for i < NMAX {
		fmt.Scan(&A[i])
		if A[i] == 0 {
			break
		}
		i++
		*n++
	}
}

func cetak(A tabInt, n int) {
	for i := 0; i < n; i++ {
		if A[i] >= 0 {
			fmt.Print(A[i], " ")
		}
	}
	fmt.Println()
}

func jumlah(A tabInt, n int) int {
	var jumlah int
	for i := 0; i < n; i++ {
		if A[i] < 0 {
			A[i] *= -1
		}
		jumlah += A[i]
	}
	return jumlah
}

func rata_rata(A tabInt, n int) float64 {
	var rata float64
	var jumlah int
	for i := 0; i < n; i++ {
		if A[i] < 0 {
			A[i] *= -1
		}
		jumlah += A[i]
	}
	rata = float64(jumlah) / float64(n)
	return rata
}
