package main

import "fmt"

func main() {
	var bil, i, j int
	fmt.Scan(&bil)
	j = 0
	for i = 1; i <= bil; i++ {
		if bil%i == 0 {
			j += 1
		}
	}
	if j == 2 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
