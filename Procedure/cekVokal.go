package main

import "fmt"

func cekHuruf(kata string) {
	/*
	   IS : Terdefinisi suatu kata bertipe string
	   FS : Menampilkan hasil apakah kata tersebut memiliki huruf vokal diawal atau diakhir.
	*/
	fmt.Scan(&kata)

	aAwal := string(kata[0]) == "a"
	iAwal := string(kata[0]) == "i"
	uAwal := string(kata[0]) == "u"
	eAwal := string(kata[0]) == "e"
	oAwal := string(kata[0]) == "o"
	vokalAwal := aAwal || iAwal || uAwal || eAwal || oAwal

	aAkhir := string(kata[len(kata)-1]) == "a"
	iAkhir := string(kata[len(kata)-1]) == "i"
	uAkhir := string(kata[len(kata)-1]) == "u"
	eAkhir := string(kata[len(kata)-1]) == "e"
	oAkhir := string(kata[len(kata)-1]) == "o"
	vokalAkhir := aAkhir || iAkhir || uAkhir || eAkhir || oAkhir

	if vokalAwal && vokalAkhir {
		fmt.Println("vokal di awal dan di akhir")
	} else if vokalAwal {
		fmt.Println("vokal di awal")
	} else if vokalAkhir {
		fmt.Println("vokal di akhir")
	} else {
		fmt.Println("tidak ada vokal")
	}
}

func main() {
	var kata string
	cekHuruf(kata)
}
