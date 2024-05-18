package main

import "fmt"

func main() {
	var u1, u2, u3 int
	fmt.Scan(&u1, &u2, &u3)
	barisAritmatika(u1, u2, u3)
}

func barisAritmatika(u1, u2, u3 int) {
	/*
		I.S. terdefinisi tiga buah bilangan bulat.
		F.S. tercetak karakter 'y' jika termasuk barisan aritmatika atau 't' jika tidak termasuk.
	*/

	if u2-u1 == u3-u2 {
		fmt.Println("ya")
	} else {
		fmt.Println("tidak")
	}
}
