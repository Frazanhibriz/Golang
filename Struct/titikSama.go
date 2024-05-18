package main

import (
	"fmt"
	"math"
)

type titik struct {
	x, y float64
}
type garis struct {
	t1, t2 titik
}

func main() {
	var g1, g2 garis
	var p1, p2 float64
	fmt.Scan(&g1.t1.x, &g1.t1.y, &g1.t2.x, &g1.t2.y, &g2.t1.x, &g2.t1.y, &g2.t2.x, &g2.t2.y)
	p1 = math.Sqrt((g1.t1.x-g1.t2.x)*(g1.t1.x-g1.t2.x) + (g1.t1.y-g1.t2.y)*(g1.t1.y-g1.t2.y))
	p2 = math.Sqrt((g2.t1.x-g2.t2.x)*(g2.t1.x-g2.t2.x) + (g2.t1.y-g2.t2.y)*(g2.t1.y-g2.t2.y))
	fmt.Println(p1 == p2)
}
