package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(coba(n))
}

func coba(n int) int {
	if n == 1 {
		return 1
	} else if n%2 == 0 {
		return n + coba(n-1)
	} else if n%2 != 0 {
		return n - coba(n-1)
	}
	return coba(n - 1)
}
