package main

import "fmt"

func main() {
	var n, apel, mangga, jeruk int
	var buah string
	fmt.Scan(&n)
	jumBuah(n, 1, buah, &apel, &mangga, &jeruk)
	fmt.Println(apel, mangga, jeruk)
}

func jumBuah(n, i int, buah string, apel, mangga, jeruk *int) {
	fmt.Scan(&buah)
	if n == i {
		if buah == "apel" {
			*apel++
		} else if buah == "mangga" {
			*mangga++
		} else if buah == "jeruk" {
			*jeruk++
		}
	} else {
		if buah == "apel" {
			*apel++
		} else if buah == "mangga" {
			*mangga++
		} else if buah == "jeruk" {
			*jeruk++
		}
		jumBuah(n, i+1, buah, apel, mangga, jeruk)
	}
}
