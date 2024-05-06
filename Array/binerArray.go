package main

import "fmt"

const nMax int = 1000

type tabInt [nMax]int

func isiArray(arr1, arr2 *tabInt, n int) {
	var i int // variabel digunakan dalam perulangan dan menyimpan sementara nilai untuk masukkan
	i = 1
	for i <= n { // mengisi array 1
		fmt.Scan(&arr1[i])
		for arr1[i] != 0 && arr1[i] != 1 {
			fmt.Scan(&arr1[i])
		}
		i++
	}
	i = 1
	for i <= n { // mengisi array 2
		fmt.Scan(&arr2[i])
		for arr2[i] != 0 && arr2[i] != 1 {
			fmt.Scan(&arr2[i])
		}
		i++
	}
}

func and(x1, x2 int) int {
	if x1 == 1 && x2 == 1 {
		return 1
	} else {
		return 0
	}
}

func or(x1, x2 int) int {
	if x1 == 1 || x2 == 1 {
		return 1
	} else if x1 == 0 || x2 == 0 {
		return 0
	} else {
		return 0
	}
}

func hitungOR(arr1, arr2 tabInt, n int) int {
	var totalOR int // variabel digunakan untuk iterasi loop dan menyimpan total banyaknya operasi yang memenuhi
	for i := 1; i <= n; i++ {
		if or(arr1[i], arr2[i]) == 1 {
			totalOR++
		}
	}
	return totalOR
}

func hitungAND(arr1, arr2 tabInt, n int) int {
	var totalAND int
	for i := 1; i <= n; i++ {
		if and(arr1[i], arr2[i]) != 0 {
			totalAND++
		}
	}
	return totalAND
}

func xor(x1, x2 int) int {
	if x1 != x2 {
		return 1
	} else {
		return 0
	}
}

func hitungXOR(arr1, arr2 tabInt, n int) int {
	var totalXOR int
	for i := 1; i <= n; i++ {
		if xor(arr1[i], arr2[i]) == 1 {
			totalXOR++
		}
	}
	return totalXOR
}

func cetak(arr1, arr2 tabInt, n int) {
	var i, totalOR, totalAND, totalXOR int
	totalOR = hitungOR(arr1, arr2, n)
	fmt.Println("Total OR:", totalOR) // menampilkan total OR yang memenuhi
	for i = 1; i <= n; i++ {
		fmt.Print(or(arr1[i], arr2[i]), " ") // menampilkan hasil hasil dari operasi logika OR(perhatikan contoh !!!)
	}
	fmt.Println()

	totalAND = hitungAND(arr1, arr2, n)
	fmt.Println("Total AND:", totalAND) // menampilkan total operasi logika XOR yang memenuhi
	for i = 1; i <= n; i++ {
		fmt.Print(and(arr1[i], arr2[i]), " ") // menampilkan hasil hasil dari operasi logika XOR(perhatikan contoh !!!)
	}
	fmt.Println()

	totalXOR = hitungXOR(arr1, arr2, n)
	fmt.Println("Total XOR:", totalXOR) // menampilkan total operasi logika XOR yang memenuhi
	for i = 1; i <= n; i++ {
		fmt.Print(xor(arr1[i], arr2[i]), " ") // menampilkan hasil hasil dari operasi logika XOR(perhatikan contoh !!!)
	}
	fmt.Println()
}

func main() {
	var n int
	fmt.Scan(&n) // membaca nilai n
	var arr1, arr2 tabInt
	isiArray(&arr1, &arr2, n)
	cetak(arr1, arr2, n)
}
