package main

import (
	"fmt"
)

const N int = 1000

func main() {
	var T [N]int
	var jumlah, target int
	fmt.Scanln(&jumlah)
	for i := 0; i < jumlah; i++ {
		fmt.Scan(&T[i])
	}
	fmt.Scan(&target)
	fmt.Println(lebihBesar(T, jumlah, target))
}

func lebihBesar(T [N]int, jumlah, target int) int {
	var banyakTarget int

	for i := 0; i < jumlah; i++ {
		if T[i] > target {
			banyakTarget++
		}
	}
	return banyakTarget
}
