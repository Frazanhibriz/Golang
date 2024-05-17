// program untuk mengecek apakah himpunan yang di masukkan dalam array memiliki anggota yang sama //

package main

import (
	"fmt"
)

const NMAX int = 100

type tabInt [NMAX]int

func main() {
	var bil tabInt
	var n int // n merupakan banyak bilangan yang akan diinputkan

	fmt.Print("Masukkan jumlah bilangan: ")
	fmt.Scan(&n)

	inputBilangan(&bil, n)
	frekuensi(bil, n)
}

func inputBilangan(b *tabInt, n int) {
	fmt.Println("Masukkan bilangan:")
	for i := 0; i < n; i++ {
		fmt.Scan(&b[i])
	}
}

func frekuensi(b tabInt, n int) {
	var found bool

	for i := 0; i < n; i++ {
		x := b[i]
		found = false
		for j := i + 1; j < n; j++ {
			if b[j] == x {
				found = true
				break
			}
		}
		if found {
			fmt.Println("Tidak Valid")
			break
		}
	}

	if !found {
		fmt.Println("Valid")
	}
}
