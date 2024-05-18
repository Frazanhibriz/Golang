package main

import "fmt"

func sudutSegitiga(a int, b int, c int) bool {
	var nilai bool

	if a <= 180 && b <= 180 && c <= 180 {
		nilai = a+b+c == 180
	}

	return nilai
}
func main() {
	var a, b, c int

	fmt.Scan(&a, &b, &c)
	fmt.Println(sudutSegitiga(a, b, c))
}
