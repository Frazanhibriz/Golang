package main

import (
	"fmt"
	"math"
)

func main() {
	type titik struct {
		x, y float64
	}
	var A, B, C titik

	fmt.Scan(&A.x, &A.y)
	fmt.Scan(&B.x, &B.y)
	fmt.Scan(&C.x, &C.y)

	radius := math.Sqrt((B.x-A.x)*(B.x-A.x) + (B.y-A.y)*(B.y-A.y))
	distance := math.Sqrt((C.x-A.x)*(C.x-A.x) + (C.y-A.y)*(C.y-A.y))

	if radius > distance {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
