package main

import "fmt"

func descending(a, b, c int) {
	var max, mid, min int

	max = a

	if max < b {
		max = b
	}

	if max < c {
		max = c
	}

	min = a

	if min > b {
		min = b
	}

	if min > c {
		min = c
	}

	mid = a + b + c - max - min

	fmt.Println(max, mid, min)
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	descending(a, b, c)
}
