package main

import (
	"fmt"
	"math"
)

type titik struct {
	x, y float64
}

func main() {
	var p1, p2 titik

	fmt.Scan(&p1.x, &p1.y, &p2.x, &p2.y)

	fmt.Println(jarak(p1, p2))
}

func jarak(p1, p2 titik) float64 {
	x := math.Pow((float64(p1.x)-float64(p2.x)), 2) + math.Pow((float64(p1.x)-float64(p2.x)), 2)
	return akar(x)
}

func akar(x float64) float64 {
	return math.Sqrt(x)
}
