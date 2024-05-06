package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	hasil := fibonacci(n)
	fmt.Println(hasil)
}

func fibonacci(n int) int {
	if n == 1 {
		return 0
	} else if n == 2 {
		return 1
	} else if n == 3 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}
