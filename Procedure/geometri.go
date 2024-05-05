package main

import "fmt"

func barisGeometri(u1, u2, u3 int) {
	/*
	   I.S. terdefinisi tiga buah bilangan bulat yang sudah terurut membesar.
	   F.S. tercetak string "ya" jika termasuk barisan geometri atau "tidak" jika tidak termasuk.
	*/
	if u2/u1 == u3/u2 {
		fmt.Println("ya")
	} else {
		fmt.Println("tidak")
	}
}

func main() {
	var u1, u2, u3 int
	fmt.Scan(&u1, &u2, &u3)
	barisGeometri(u1, u2, u3)
}
