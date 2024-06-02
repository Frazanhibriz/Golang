package main

import (
	"fmt"
)

func signupPasien() {
	p := &DataPasien{}

	fmt.Print("Masukkan nama lengkap : ")
	fmt.Scan(&p.NamaLengkap)

	fmt.Print("Masukkan email : ")
	fmt.Scan(&p.Email)

	fmt.Print("Masukkan password : ")
	fmt.Scan(&p.Password)

	fmt.Print("Masukkan tanggal lahir : ")
	fmt.Scan(&p.Lahir.Tanggal)

	fmt.Print("Masukkan bulan lahir : ")
	fmt.Scan(&p.Lahir.Bulan)

	fmt.Print("Masukkan tahun lahir : ")
	fmt.Scan(&p.Lahir.Tahun)

	fmt.Print("Masukkan kota lahir : ")
	fmt.Scan(&p.Lahir.Kota)

	if cekSudahDaftarPasien(p) {
		fmt.Println("Akun sudah terdaftar")
	} else {
		DaftarPasien[jumlahPasien] = p
		jumlahPasien++

		fmt.Println("Pendaftaran berhasil")
	}
}

func cekSudahDaftarPasien(p *DataPasien) bool {
	for i := 0; i < jumlahPasien; i++ {
		if DaftarPasien[i].Email == p.Email {
			return true
		}
	}
	return false
}

func loginPasien() {
	var loginEmail, loginPassword string

	fmt.Print("Masukkan email Anda: ")
	fmt.Scan(&loginEmail)

	fmt.Print("Masukkan password Anda: ")
	fmt.Scan(&loginPassword)

	if cekLoginPasien(loginEmail, loginPassword) {
		fmt.Println("Login berhasil")
		forumKonsultasiPasien()
	} else {
		fmt.Println("Login gagal")
	}
}

func cekLoginPasien(loginEmail, loginPassword string) bool {
	for i := 0; i < jumlahPasien; i++ {
		if DaftarPasien[i].Email == loginEmail && DaftarPasien[i].Password == loginPassword {
			pasienSedangLogin = DaftarPasien[i]
			return true
		}
	}
	return false
}
