package main

import "fmt"

func main() {
	var x, n int
	var hasil bool
	fmt.Scan(&x, &n)
	hasil = false
	for n > 0 {
		if n%10 == x {
			hasil = true
		}
		n = n / 10
	}
	fmt.Println(hasil)
}
