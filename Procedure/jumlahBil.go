package main

import "fmt"

func main() {
	var x, y, hasil int

	fmt.Scan(&x, &y)

	for x != 0 && y != 0 {
		jumlah(x, y, &hasil)
		fmt.Scan(&x, &y)
	}
	fmt.Println(hasil)
}

func jumlah(x, y int, hasil *int) {
	if x%2 != 0 && y%2 == 0 {
		*hasil = *hasil + x + y
	}
}
