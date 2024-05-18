package main

import "fmt"

func main() {
	var n, r int
	fmt.Scan(&n, &r)
	fmt.Println(permutasi(n, r))
}

func permutasi(n, r int) int {
	if n == (n-r) || n == 1 {
		return 1
	} else {
		return n * permutasi(n-1, r-1)
	}
}
