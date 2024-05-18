package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	fmt.Print(sumRepeatUntillTen(N))
}

func sumRepeatUntillTen(N int) int {
	jumlah := 0
	if N == 0 {
		jumlah = 1
	}
	for i := 1; i <= (10 * N); i++ {
		jumlah += i
	}
	return jumlah
}
