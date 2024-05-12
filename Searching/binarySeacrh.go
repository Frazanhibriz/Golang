package main

import "fmt"

const NMAX int = 10

type tabInt [NMAX]int

func main() {
	var data tabInt
	var nData int
	var x2 int
	var idx int

	fmt.Scan(&x2)

	bacaData(&data, &nData)

	cetakData(data, nData)

	fmt.Print("Hasil pencarian: ")
	binarySearch(data, nData, x2, &idx)

	if idx != -1 {
		fmt.Printf("Bilangan %d ditemukan pada posisi ke-%d\n", x2, idx+1)
	} else {
		fmt.Println("Bilangan tidak ditemukan.")
	}
}

func bacaData(A *tabInt, n *int) {
	for {
		if *n >= NMAX {
			break
		}
		fmt.Scan(&A[*n])
		if *n > 0 && A[*n] < A[*n-1] {
			break
		}
		*n++
	}
}

func cetakData(A tabInt, n int) {
	fmt.Print("Data bilangan: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", A[i])
	}
	fmt.Println()
}

func binarySearch(A tabInt, n int, x int, idx *int) {
	var left, right, mid int
	left = 0
	right = n - 1
	mid = (left + right) / 2

	for left <= right && A[mid] != x {
		if x < A[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
		mid = (left + right) / 2
	}
	if left <= right {
		*idx = mid
	} else {
		*idx = -1
	}
}
