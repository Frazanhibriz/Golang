package main

import "fmt"

func konversiMeterMil(meter float64) float64 {
	// fungsi mengembalikan nilai mil dengan masukan meter
	var mil float64
	mil = meter * 0.00062137119
	return mil
}

func main() {
	var meter, mil float64
	fmt.Scan(&meter)
	mil = konversiMeterMil(meter)
	fmt.Println(mil)
}
