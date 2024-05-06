package main

import "fmt"

type SkorGame struct {
	TotalA   int
	TotalB   int
	TotalC   int
	max      int
	Pemenang byte
}

var skor SkorGame

func isGenap(n int) bool {
	return n%2 == 0
}

func maxSkor() {
	skor.max = skor.TotalA

	if skor.max < skor.TotalB {
		skor.max = skor.TotalB
		skor.Pemenang = 'B'
	}

	if skor.max < skor.TotalC {
		skor.max = skor.TotalC
		skor.Pemenang = 'C'
	}

	if skor.max == skor.TotalA {
		skor.Pemenang = 'A'
	}

}

func gameDadu(n int) {
	var d1, d2, d3 int

	for i := 0; i < n; i++ {
		fmt.Scan(&d1, &d2, &d3)

		if isGenap(d1) {
			skor.TotalA += d1
		}
		if isGenap(d2) {
			skor.TotalB += d2
		}
		if isGenap(d3) {
			skor.TotalC += d3
		}
	}
	maxSkor()
}

func main() {
	var n int

	fmt.Scan(&n)

	gameDadu(n)
	fmt.Printf("%c %d\n", skor.Pemenang, skor.max)

}
