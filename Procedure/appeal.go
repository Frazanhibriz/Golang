package main

import "fmt"

// buat prosedurnya saja

func lebihKecil(b1, b2, banding int) {
	/*
		I.S. terdefinisi dua buah bilangan bulat yang berbeda.
		F.S. tercetak bilangan masukan yang lebih kecil.
	*/

	if b1 < b2 {
		banding = b1
	} else if b1 > b2 {
		banding = b2
	}
	fmt.Println(banding)
}

func main() {
	var b1, b2 int
	var banding int

	fmt.Scan(&b1, &b2)

	lebihKecil(b1, b2, banding)
}
