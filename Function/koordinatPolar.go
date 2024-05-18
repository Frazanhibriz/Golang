package main

import (
	"fmt"
	"math"
)

func panjangX(r float64, t float64) float64 {
	var x float64

	t = t * math.Pi / 180
	x = r * math.Cos(t)

	return x
}

func panjangY(r float64, t float64) float64 {
	var x float64

	t = t * math.Pi / 180
	x = r * math.Sin(t)

	return x
}

func main() {
	var r, t float64

	fmt.Scanf("%f %f", &r, &t)
	fmt.Printf("%.2f %.2f\n", panjangX(r, t), panjangY(r, t))
}
