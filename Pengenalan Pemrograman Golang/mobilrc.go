package main

import "fmt"

func main() {
	var tim, pemenang string
	var timA, timB int
	for i := 1; i <= 10; i++ {
		fmt.Scan(&tim)
		if tim == "A" {
			timA++
		} else if tim == "B" {
			timB++
		}

		if pemenang == "" {
			if timA == 4 {
				pemenang = "A"
			} else if timB == 4 {
				pemenang == "B"
			}
		}
	}
	if timA != 5 || timB != 5 {
		pemenang == "tidak valid"
	}
	fmt.Println(pemenang)
}