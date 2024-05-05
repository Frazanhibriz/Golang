package main

import "fmt"

func main() {
	var n, r int
	fmt.Scan(&n, &r)
	permutasi(n, r)
}

func faktorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}

func permutasi(x, y int) {
	var n, m, z, hasilPer int
	if x >= y {
		z = x - y
		faktorial(x, &n)
		faktorial(z, &m)
	} else if x < y {
		z = y - x
		faktorial(y, &n)
		faktorial(z, &m)
	}
	hasilPer = n / m
	fmt.Println(hasilPer)
}
