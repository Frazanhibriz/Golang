package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	emas(n)
}

func emas(n int) {
	var anggaran, adminCost, emas10, emas5, emas1 int
	for i := 0; i < n; i++ {
		fmt.Scan(&anggaran)
		for anggaran >= 9000 {
			emas10++
			anggaran -= 9000
			adminCost += 1000
		}

		for anggaran >= 5000 {
			emas5++
			anggaran -= 5000
			adminCost += 3000
		}

		for anggaran >= 1500 {
			emas1++
			anggaran -= 1500
			adminCost += 5000
		}
	}
	fmt.Println(emas10, emas5, emas1, adminCost)
}
