package main

import "fmt"

func signupDokter() {
	d := &DataDokter{}

	fmt.Println("Masukkan nama lengkap : ")
	fmt.Scan(&d.NamaLengkap)

	fmt.Println("Masukkan email : ")
	fmt.Scan(&d.Email)

	fmt.Println("Masukkan password : ")
	fmt.Scan(&d.Password)

	fmt.Println("Masukkan tanggal lahir : ")
	fmt.Scan(&d.Lahir.Tanggal)

	fmt.Println("Masukkan bulan lahir : ")
	fmt.Scan(&d.Lahir.Bulan)

	fmt.Println("Masukkan tahun lahir : ")
	fmt.Scan(&d.Lahir.Tahun)

	fmt.Println("Masukkan kota lahir : ")
	fmt.Scan(&d.Lahir.Kota)

	fmt.Println("Masukkan spesialisasi anda : ")
	fmt.Scan(&d.Spesialisasi)

	if cekSudahDaftarDokter(d) {
		fmt.Println("Akun sudah terdaftar")
	} else {
		DaftarDokter[jumlahDokter] = d
		jumlahDokter++
		fmt.Println("Pendaftaran berhasil")
	}
}

func cekSudahDaftarDokter(d *DataDokter) bool {
	for i := 0; i < jumlahDokter; i++ {
		if DaftarDokter[i].Email == d.Email {
			return true
		}
	}
	return false
}

func loginDokter() {
	var loginEmail, loginPassword string

	fmt.Println("Masukkan email anda :")
	fmt.Scan(&loginEmail)

	fmt.Println("Masukkan password anda :")
	fmt.Scan(&loginPassword)

	if cekLoginDokter(loginEmail, loginPassword) {
		fmt.Println("Login berhasil")
		forumKonsultasiDokter()
	} else {
		fmt.Println("Login gagal")
	}
}

func cekLoginDokter(loginEmail, loginPassword string) bool {
	for i := 0; i < jumlahDokter; i++ {
		if DaftarDokter[i].Email == loginEmail && DaftarDokter[i].Password == loginPassword {
			return true
		}
	}
	return false
}

func cariSpesialisasi() {
	var spesialisasi string

	fmt.Print("Masukkan spesialisasi yang dicari : ")
	fmt.Scan(&spesialisasi)

	tampilkanSemuaDokter("", spesialisasi)
}

func cariDokter() {
	var dokter string

	fmt.Print("Masukkan nama dokter yang dicari : ")
	fmt.Scan(&dokter)

	tampilkanSemuaDokter(dokter, "")
}

func tampilkanSemuaDokter(namaDokter, spesialisasi string) {
	var pilih string
	var daftarDokter [NMAX]*DataDokter
	var idxDokter, dokter int

	fmt.Println("--------------------------------------------------")
	fmt.Println("|                  DAFTAR DOKTER                 |")
	fmt.Println("--------------------------------------------------")

	for i := 0; i < jumlahDokter; i++ {
		if namaDokter == "" || spesialisasi == "" ||
			(namaDokter != "" && namaDokter == DaftarDokter[i].NamaLengkap) ||
			(spesialisasi != "" && spesialisasi == DaftarDokter[i].Spesialisasi) {

			daftarDokter[idxDokter] = DaftarDokter[i]
			idxDokter++

			fmt.Printf("%d. Nama: %s , Spesialisasi: %s \n", idxDokter, DaftarDokter[i].NamaLengkap, DaftarDokter[i].Spesialisasi)
		}
	}

	fmt.Println("--------------------------------------------------")
	fmt.Print("Ingin kirim pertanyaan dengan dokter pilihan anda? Y/N ")
	fmt.Scan(&pilih)

	if pilih == "Y" || pilih == "y" {
		fmt.Print("Masukkan nomor dokter pilihan anda : ")
		fmt.Scan(&dokter)

		postPertanyaan(daftarDokter[dokter-1])
	}

	return
}
