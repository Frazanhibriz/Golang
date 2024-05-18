package main

import "fmt"

func main() {
	var golduser, platinumuser, silveruser int
	var total, point, goldtotal, platinumtotal, silvertotal, hasil1, hasil2, hasil3 float64

	for total < 500 {
		fmt.Scan(&point)
		for point < 50 {
			fmt.Scan(&point)
		}

		if point >= 50 {
			total += point
			if point > 200 {
				goldtotal += point
				golduser++
			} else if point >= 100 && point <= 200 {
				platinumtotal += point
				platinumuser++
			} else if point >= 50 && point <= 99 {
				silvertotal += point
				silveruser++
			}
		}
	}
	hasil1 = goldtotal / float64(golduser)
	hasil2 = platinumtotal / float64(platinumuser)
	hasil3 = silvertotal / float64(silveruser)
	if hasil1 > 0 {
		fmt.Printf("Rata-rata Gold User: %.2f\n", hasil1)
	} else {
		fmt.Println("Rata-rata Gold User:", 0)
	}
	if hasil2 > 0 {
		fmt.Printf("Rata-rata Platinum User: %.2f\n", hasil2)
	} else {
		fmt.Println("Rata-rata Platinum User:", 0)
	}
	if hasil3 > 0 {
		fmt.Printf("Rata-rata Silver User: %.2f\n", hasil3)
	} else {
		fmt.Println("Rata-rata Silver User:", 0)
	}

}
