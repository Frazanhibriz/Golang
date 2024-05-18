package main

import "fmt"

func main() {
	var i, x, masukan int
	fmt.Scan(&masukan)
	for i = 1; i <= masukan; i++ {
		for x = 1; x <= masukan; x++ {
			fmt.Print(i, " ")
		}
		fmt.Println()
	}
}
