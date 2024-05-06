package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	coba(n, 1)
}

func coba(n, i int) {
	if n == i {
		if i%2 != 0 {
			fmt.Println(i, " ")
		}
	} else {
		if i%2 != 0 {
			fmt.Println(i, " ")
		}
		coba(n, i+1)
	}
}
