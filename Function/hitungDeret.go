package main

import "fmt"

func hitungDeret(n int, m int) float64 {
	var jumlahDeret float64 = 0
	// var temp int = n
	for i := n; i <= m; i += 2 {
		jumlahDeret += 3 / float64(i)
	}
	for i := n + 1; i <= m; i += 2 {
		jumlahDeret -= 3 / float64(i)
	}
	return jumlahDeret
}
func main() {
	var n, m int

	fmt.Scan(&n, &m)
	fmt.Printf("%.2f\n", hitungDeret(n, m))
}
