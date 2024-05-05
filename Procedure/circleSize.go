package main

import "fmt"

func hitungLuasKelilingLingkaran(r float64, l *float64, k *float64) {
	const pi = 3.14
	*l = pi * r * r
	*k = 2 * pi * r
}

func hitungLuasKelilingPersegi(s float64, l *float64, k *float64) {
	*l = s * s
	*k = 4 * s
}

func hitungTotal(lL float64, lP float64, kL float64, kP float64, toLuas *float64, toKeliling *float64) {
	*toLuas = lL + lP
	*toKeliling = kL + kP
}

func main() {
	var r, s, lL, lP, kL, kP, toLuas, toKeliling float64
	fmt.Scan(&r, &s)

	for r != 0 && s != 0 {

		hitungLuasKelilingLingkaran(r, &lL, &kL)
		hitungLuasKelilingPersegi(s, &lP, &kP)
		hitungTotal(lL, lP, kL, kP, &toLuas, &toKeliling)

		fmt.Printf("%7s %7s %7s %7s %7s %7s %7s %7s\n", "R", "S", "LL", "LP", "KL", "KP", "TL", "TP")
		fmt.Printf("%7.2f %7.2f %7.2f %7.2f %7.2f %7.2f %7.2f %7.2f\n", r, s, lL, lP, kP, kL, toLuas, toKeliling)
		fmt.Scan(&r, &s)
	}
}
