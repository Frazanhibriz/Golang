package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(Des2Bin(n))
}

func Des2Bin(desimal int) string {
	var n, sisa int
	var biner string
	if desimal == 0 {
		return "0"
	}
	for desimal > 0 {
		divison(desimal, 2, &n, &sisa)
		biner = Num2Str(sisa) + biner
		desimal = n
	}
	return biner
}

func divison(a, b int, result, remainder *int) {
	*result = a / b
	*remainder = a % b
}

func Num2Str(x int) string {
	var hasil string
	if x == 0 {
		hasil = "0"
	} else if x == 1 {
		hasil = "1"
	}
	return hasil
}
