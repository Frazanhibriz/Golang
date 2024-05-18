package main

import "fmt"

type pegawai struct {
	nama                    string
	masa_kerja              int
	gaji_pokok, besar_bonus float64
}

func main() {
	var p pegawai

	fmt.Scan(&p.nama, &p.gaji_pokok, &p.masa_kerja)

	hitung_bonus(&p)

	fmt.Printf("Pegawai dengan nama %s mendapatkan bonus sebesar Rp. %0.f\n", p.nama, p.besar_bonus)
}

func hitung_bonus(p *pegawai) {
	if p.masa_kerja >= 10 {
		*&p.besar_bonus = p.gaji_pokok * 1.5
	} else if p.masa_kerja >= 5 && p.masa_kerja < 10 {
		*&p.besar_bonus = p.gaji_pokok * 0.75
	} else if p.masa_kerja < 5 {
		*&p.besar_bonus = p.gaji_pokok * 0.5
	}
}
