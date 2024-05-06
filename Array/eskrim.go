package main

import "fmt"

type arrayKata [8]string

func main() {
	var arrKata arrayKata
	isiArray(&arrKata)
	fmt.Println(arrKata, "diubah menjadi :", arraytostring(arrKata))
}

func isiArray(arrKata *arrayKata) {
	/*I.S. Terdefinisi array untuk menampung string huruf
	  F.S. arrKata berisi 8 huruf unik*/
	for i := 0; i < 8; i++ {
		fmt.Scan(&arrKata[i])
	}

}

func arraytostring(arrKata arrayKata) string {
	/* mengembalikan array hasil konversi 8 huruf menjadi sebuah string */
	var hasil string
	for i := 0; i < 8; i++ {
		hasil += string(arrKata[i])
	}
	return hasil
}
