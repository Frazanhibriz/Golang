package main

import "fmt"

func main() {
	type pemain struct {
		nama        string
		gol, assist int
	}
	var p1, p2, p3 pemain

	fmt.Scan(&p1.nama, &p1.gol, &p1.assist)
	fmt.Scan(&p2.nama, &p2.gol, &p2.assist)
	fmt.Scan(&p3.nama, &p3.gol, &p3.assist)

	if p1.gol < p2.gol && p1.gol < p3.gol {
		fmt.Println(p1.nama)
	} else if p2.gol < p1.gol && p2.gol < p3.gol {
		fmt.Println(p2.nama)
	} else {
		fmt.Println(p3.nama)
	}

	if p1.assist < p2.assist && p1.assist < p3.assist {
		fmt.Println(p1.nama)
	} else if p2.assist < p1.assist && p2.assist < p3.assist {
		fmt.Println(p2.nama)
	} else {
		fmt.Println(p3.nama)
	}
}
