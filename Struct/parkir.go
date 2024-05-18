package main

import "fmt"

func totalMotor(n int) int {
	// Mengembalikan total besaran jika tipe kendaraan adalah sepeda motor.
	total := 0
	for i := 0; i < n; i++ {
		if parkir[i].tipe == 'm' {
			total += parkir[i].besaran
		}
	}
	return total
}
func totalMobil(n int) int {
	// Mengembalikan total besaran jika tipe kendaraan adalah mobil.
	total := 0
	for i := 0; i < n; i++ {
		if parkir[i].tipe == 'b' {
			total += parkir[i].besaran
		}
	}
	return total
}

type Parkir struct {
	tipe    byte
	besaran int
}

const NMax int = 1000

var parkir [NMax]Parkir

func main() {
	/* I.S.: Terdefinisi nilai bilangan bulat n. Data string sejumlah n buah tersedia
	   pada perangkat input. */
	/* F.S.: Array pendapatan parkir diisi dengan data yang diberikan. */

	var n, i int

	fmt.Scan(&n)
	for i = 0; i < n; i++ {

		//fmt.Scan(&parkir[i].tipe)
		//fmt.Scan(&parkir[i].besaran)

		for parkir[i].tipe != 'm' && parkir[i].tipe != 'b' {
			fmt.Scanf("%c", &parkir[i].tipe)
			fmt.Scanf("%d", &parkir[i].besaran)
		}
	}
	fmt.Println(totalMobil(n), totalMotor(n))
}
