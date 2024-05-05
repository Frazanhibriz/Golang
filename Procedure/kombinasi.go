package main

import "fmt"

func main() {
	var n, r int
	fmt.Scan(&n, &r)
	kombinasi(n, r)
}

func faktorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil *= i
	}
}

func kombinasi(x, y int) {
	var n, m, o, z, hasilKom int
	if x >= y {
		z = x - y
		faktorial(x, &n)
		faktorial(y, &m)
		faktorial(z, &o)
	} else if x < y {
		z = y - x
		faktorial(y, &n)
		faktorial(x, &m)
		faktorial(z, &o)
	}
	hasilKom = n / (m * o)
	fmt.Println(hasilKom)
}
