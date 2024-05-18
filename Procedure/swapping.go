package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	for a != b {
		swap(&a, &b)
		fmt.Println(a, b)
		fmt.Scan(&a, &b)
	}
}

func swap(a, b *int) {
	if *a > *b {
		*a, *b = *b, *a
	}
}
