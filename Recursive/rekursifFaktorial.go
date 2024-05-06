package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	hasil := faktorial(x)
	fmt.Println(hasil)
}

func faktorial(x int) int {
	if x == 1 {
		return 1
	} else {
		return x * faktorial(x-1)
	}
}
