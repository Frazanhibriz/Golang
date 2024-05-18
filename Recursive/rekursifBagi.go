package main

import "fmt"

func main() {
	var m, n int
	fmt.Scan(&m, &n)
	fmt.Println(div(m, n, 0))
}

func div(m, n, i int) int {
	if m < n {
		return i
	}
	i++
	return div(m-n, n, i)
}
