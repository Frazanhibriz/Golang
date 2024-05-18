package main

import "fmt"

func main() {
	var bil, max int
	fmt.Scan(&bil)
	max = bil % 10
	for bil > 0 {
		bil = bil / 10
		if bil%10 > max {
			max = bil % 10
		}
	}
	fmt.Println(max)
}
