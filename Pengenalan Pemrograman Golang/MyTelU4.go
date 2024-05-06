package main

import "fmt"

func main() {
	var total, point, golduser, platinumuser, silveruser int

	for total < 500 {
		fmt.Scan(&point)
		for point < 50 {
			fmt.Scan(&point)
		}

		total += point

		if point > 200 {
			golduser++
		} else if point >= 100 && point <= 200 {
			platinumuser++
		} else if point >= 50 && point <= 99 {
			silveruser++
		}
	}

	fmt.Println("Jumlah Gold User :", golduser)
	fmt.Println("Jumlah Platinum User :", platinumuser)
	fmt.Println("Jumlah Silver User :", silveruser)
}
