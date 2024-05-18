package main

import "fmt"

func main() {
	var x, y int

	fmt.Scan(&x, &y)

	for x != y {
		mengurut(&x, &y)
		fmt.Println(x, y)
		fmt.Scan(&x, &y)
	}

}

func mengurut(a *int, b *int) {
	if *a > *b {
		*a, *b = *b, *a
	}
}
