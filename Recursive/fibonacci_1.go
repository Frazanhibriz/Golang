package main

import "fmt"

func main() {
	var bil int
	fmt.Scan(&bil)
	hasil := fibonacci(bil)
	fmt.Println(hasil)
}

func fibonacci(bil int) int {
	if bil <= 3 {
		return 1
	} else {
		return fibonacci(bil-1) + fibonacci(bil-2) + fibonacci(bil-3)
	}
}
