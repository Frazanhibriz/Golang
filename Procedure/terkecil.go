package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Scan(&a, &b, &c)

	terkecil(a, b, c)
}

func terkecil(a, b, c int) {
	/* I.S. terdefinisi 3 bilangan bulat a, b, dan c yang berbeda
	   F.S. print bilangan terkecil dari 3 bilangan bulat a,b,c */
	if a < b && a < c {
		fmt.Println(a)
	} else if b < a && b < c {
		fmt.Println(b)
	} else if c < a && c < b {
		fmt.Println(c)
	}
}
