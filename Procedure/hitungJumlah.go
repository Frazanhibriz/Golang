package main

import "fmt"

func main() {
	var h, b1, b2 int

	fmt.Scan(&b1, &b2)

	hitungJumlah(b1, b2, &h)

	fmt.Println(h)
}

func hitungJumlah(b1 int, b2 int, h *int) {
	*h = b1 + b2
}
