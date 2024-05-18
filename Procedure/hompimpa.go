package main

import "fmt"

func main() {
	var p1, p2, p3 string

	fmt.Scan(&p1, &p2, &p3)

	pemenang(p1, p2, p3)

	//fmt.Println(pemenang)
}

func pemenang(p1 string, p2 string, p3 string) {

	if p1 != p2 && p1 != p3 && p2 == p3 {
		fmt.Println("pemain 1 menang")
	} else if p2 != p1 && p2 != p3 && p1 == p3 {
		fmt.Println("pemain 2 menang")
	} else if p3 != p1 && p3 != p2 && p1 == p2 {
		fmt.Println("pemain 3 menang")
	} else if p1 == p2 && p2 == p3 {
		fmt.Println("imbang")
	}
}

