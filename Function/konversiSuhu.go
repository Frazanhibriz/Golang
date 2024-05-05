package main

import "fmt"

func reamur(C float64) float64 {
	return C * (4.0 / 5.0)
}

func fahrenheit(C float64) float64 {
	return (9.0 / 5.0 * C) + 32.0
}

func kelvin(C float64) float64 {
	return C + 273.0
}

func main() {
	var c, celciusAwal, celsiusAkhir, step float64
	var r, f, k float64

	fmt.Scan(&celciusAwal, &celsiusAkhir, &step)
	fmt.Printf("%10s %10s %10s %10s\n", "Celsius", "Reamur", "Fahrenheit", "Kelvin")

	for c = celciusAwal; c <= celsiusAkhir; c += step {
		r = reamur(c)
		f = fahrenheit(c)
		k = kelvin(c)

		fmt.Printf("%10.2f %10.2f %10.2f %10.2f\n", c, r, f, k)
	}
}
