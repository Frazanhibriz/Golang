package main

import "fmt"

func main() {
	var x, y float64

	fmt.Scan(&x, &y)

	if x > 0 && y > 0 {
		fmt.Println("kuadran I")
	} else if x < 0 && y > 0 {
		fmt.Println("kuadran II")
	} else if x < 0 && y < 0 {
		fmt.Println("kuadran III")
	} else if x > 0 && y < 0 {
		fmt.Println("kuadran IV")
	}
}
