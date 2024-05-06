package main

import "fmt"

func main() {
	var pengunjung, data, jumlah_pengunjung, before, kenaikan int
	var rata_rata float64
	fmt.Scan(&pengunjung)
	for pengunjung > 0 && pengunjung <= 200 {
		data++
		if data > 1 && before < pengunjung {
			kenaikan++
		}
		before = pengunjung
		jumlah_pengunjung = jumlah_pengunjung + pengunjung
		fmt.Scan(&pengunjung)
	}
	if data <= 0 {
		data = 1
	}
	rata_rata = float64(jumlah_pengunjung) / float64(data)
	fmt.Printf("%d %.2f", kenaikan, rata_rata)
}

// Falah Task //
