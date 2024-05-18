package main

import "fmt"

func main() {
	var i, x, masukan int
	fmt.Scan(&masukan)
	for i = 1; i <= masukan; i++ {
		for x = 1; x <= masukan; x++ {
			if i == x || x == (masukan-i)+1 {
				fmt.Print(i, " ")
			} else {
				fmt.Print("  ") // spasi 2 kali //
			}
		}
		fmt.Println()
	}
}
