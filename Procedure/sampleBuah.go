package main

import "fmt"

func main() {
	var jumBuah, jumApel, jumJeruk, jumMangga int

	fmt.Scan(&jumBuah)

	sumBuah(jumBuah, &jumApel, &jumJeruk, &jumMangga)

	fmt.Println(jumApel, jumJeruk, jumMangga)
}

func sumBuah(jumBuah int, jumApel *int, jumJeruk *int, jumMangga *int) {
	var jenisBuah string

	for i := 0; i < jumBuah; i++ {
		fmt.Scan(&jenisBuah)

		nBuah(jenisBuah, jumApel, jumJeruk, jumMangga)
	}
}

func nBuah(jenisBuah string, jumApel *int, jumJeruk *int, jumMangga *int) {
	if jenisBuah == "apel" {
		*jumApel++
	} else if jenisBuah == "jeruk" {
		*jumJeruk++
	} else if jenisBuah == "mangga" {
		*jumMangga++
	}
}
