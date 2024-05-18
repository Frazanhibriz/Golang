package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	fmt.Println(modulusLoop(n, m))
}

func modulusLoop(n, m int) int {
	var i int

	if n == 0 || m > n {
		return n
	} else if n == m {
		return 0
	}

	i = n
	for i >= m { // selama pembilang masih lebih dari atau sama dengan penyebut
		i = i - m
	}
	return i
}
