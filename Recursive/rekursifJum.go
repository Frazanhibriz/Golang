package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	hasil := jumlah(n)
	fmt.Println(hasil)
}

func jumlah(n int) int {
	if n == 0 {
		return 0
	} else {
		return n + jumlah(n-1)
	}
}
