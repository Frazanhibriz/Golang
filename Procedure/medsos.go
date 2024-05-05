package main

import "fmt"

func main() {
	var jumHari int
	fmt.Scan(&jumHari)
	teksGambarWhatsapp(jumHari)

}

func teksGambarWhatsapp(jumHari int) {
	var teks, gambar, postingaTeks, postingaGambar int

	for i := 0; i < jumHari; i++ {
		fmt.Scan(&teks, &gambar)
		postingaTeks += teks
		postingaGambar += gambar
	}
	fmt.Println(postingaTeks, postingaGambar)
}
