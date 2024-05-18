package main

import "fmt"

func main() {
	var N, hasil int
	fmt.Scan(&N)
	hasil = kakiHewan(N)
	fmt.Println(hasil)
}

func kakiHewan(N int) int {
	if N == 0 {
		return 0
	} else if N%2 != 0 {
		return 2 + kakiHewan(N-1)
	} else {
		return 3 + kakiHewan(N-1)
	}
}
