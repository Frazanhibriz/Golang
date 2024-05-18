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

	dAB := math.Sqrt(math.Pow((A.x-B.x), 2) + math.Pow((A.y-B.y), 2))
	dAC := math.Sqrt(math.Pow((A.x-C.x), 2) + math.Pow((A.y-C.y), 2))
	dBC := math.Sqrt(math.Pow((B.x-C.x), 2) + math.Pow((B.y-C.y), 2))
	fmt.Println(dAB + dBC + dAC)
}
