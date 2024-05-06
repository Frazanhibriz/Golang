package main

import "fmt"

func main() {
	var i, x, masukan int
	fmt.Scan(&masukan)
	if masukan%2 != 0 {
		for i = 1; i <= masukan; i++ {
			for x = 1; x <= masukan; x++ {
				if i == x {
					fmt.Print(" ")
				} else {
					fmt.Print("*")
				}
			}
			fmt.Println()
		}
	}
}
