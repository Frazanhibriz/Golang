package main

import "fmt"

type mobil struct {
	merek                     string
	tahun_produksi, kecepatan int
}

func main() {
	var m1, m2, m3 mobil
	var rata_rata_kecepatan float64

	fmt.Scan(&m1.merek, &m1.tahun_produksi, &m1.kecepatan)
	fmt.Scan(&m2.merek, &m2.tahun_produksi, &m2.kecepatan)
	fmt.Scan(&m3.merek, &m3.tahun_produksi, &m3.kecepatan)

	rata_rata_kecepatan = (float64(m1.kecepatan) + float64(m2.kecepatan) + float64(m3.kecepatan)) / float64(3)

	fmt.Printf("Rata rata kecepatan mobil %s (%d), ", m1.merek, m1.tahun_produksi)
	fmt.Printf("mobil %s (%d), dan mobil %s (%d): ", m2.merek, m2.tahun_produksi, m3.merek, m3.tahun_produksi)
	fmt.Printf("%.2f\n", rata_rata_kecepatan)
}
