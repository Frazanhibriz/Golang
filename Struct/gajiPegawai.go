package main

import "fmt"

type gaji struct {
	gajiPokok, pajakPenghasilan, asuransi int
}

func main() {
	var g gaji
	var total int

	fmt.Scan(&g.gajiPokok, &g.pajakPenghasilan, &g.asuransi)
	total = int(float64(g.gajiPokok) - (float64(g.gajiPokok) * (float64(g.pajakPenghasilan) + float64(g.asuransi)) / 100))
	fmt.Println(total)

}
