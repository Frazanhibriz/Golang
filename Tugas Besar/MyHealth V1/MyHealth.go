package main

import "fmt"

const NMAX int = 1000

type DataLahir struct {
	Tanggal, Bulan, Tahun int
	Kota                  string
}

type DataPasien struct {
	NamaLengkap, Email, Password string
	Lahir                        DataLahir
}

type DataDokter struct {
	NamaLengkap, Email, Password, Spesialisasi string
	Lahir                                      DataLahir
}

type DataKonsultasi struct {
	Pertanyaan string
	Pasien     DataPasien
	Tanggapan  string
	Dokter     DataDokter
	Tags       string
}

var jumlahPasien int
var jumlahDokter int
var jumlahPertanyaan int

var pasienBaru DaftarPasien
var dokterBaru DaftarDokter
var pertanyaanBaru DaftarPertanyaan

type DaftarPasien [NMAX]DataPasien
type DaftarDokter [NMAX]DataDokter
type DaftarPertanyaan [NMAX]DataKonsultasi

func main() {
	var pilih int

	fmt.Println("------------------------------")
	fmt.Println("*    Welcome to My Health    *")
	fmt.Println("------------------------------")

	for {
		menu(&pilih)

		switch pilih {
		case 1: // daftar - signup
			var profesiSignup string

			fmt.Println("Daftar sebagai pasien atau dokter?")
			fmt.Scan(&profesiSignup)

			if profesiSignup == "pasien" || profesiSignup == "Pasien" {
				signupPasien()
			} else if profesiSignup == "dokter" || profesiSignup == "Dokter" {
				signupDokter()
			}

		case 2: // masuk - login
			var profesiLogin string

			fmt.Println("Masuk sebagai pasien atau dokter?")
			fmt.Scan(&profesiLogin)

			if profesiLogin == "pasien" || profesiLogin == "Pasien" {
				loginPasien()
			} else if profesiLogin == "dokter" || profesiLogin == "Dokter" {
				loginDokter()
			}

		case 3: // konsultasi
			forumKonsultasi()

		case 4: // keluar - logout
			fmt.Println("Terima kasih telah menggunakan layanan kami.")
			return // Exit the loop and end the program
		}
	}
}

func menu(pilih *int) {
	fmt.Println("------------------")
	fmt.Println("*   MENU UTAMA   *")
	fmt.Println("------------------")
	fmt.Println("1. Sign Up")
	fmt.Println("2. Log In")
	fmt.Println("3. Forum Konsultasi")
	fmt.Println("4. Exit")
	fmt.Println("------------------")

	fmt.Println("Silahkan pilih menu : 1/2/3/4")
	fmt.Scan(pilih)
}

func signupPasien() {
	var p DataPasien

	fmt.Println("Masukkan nama lengkap : ")
	fmt.Scan(&p.NamaLengkap)

	fmt.Println("Masukkan email : ")
	fmt.Scan(&p.Email)

	fmt.Println("Masukkan password : ")
	fmt.Scan(&p.Password)

	fmt.Println("Masukkan tanggal lahir : ")
	fmt.Scan(&p.Lahir.Tanggal)

	fmt.Println("Masukkan bulan lahir : ")
	fmt.Scan(&p.Lahir.Bulan)

	fmt.Println("Masukkan tahun lahir : ")
	fmt.Scan(&p.Lahir.Tahun)

	fmt.Println("Masukkan kota lahir : ")
	fmt.Scan(&p.Lahir.Kota)

	if cekSudahDaftarPasien(p) {
		fmt.Println("Akun sudah terdaftar")
	} else {
		pasienBaru[jumlahPasien] = p
		jumlahPasien++
		fmt.Println("Pendaftaran berhasil")
	}
}

func cekSudahDaftarPasien(p DataPasien) bool {
	for i := 0; i < jumlahPasien; i++ {
		if pasienBaru[i].Email == p.Email {
			return true
		}
	}
	return false
}

func signupDokter() {
	var d DataDokter

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
		dokterBaru[jumlahDokter] = d
		jumlahDokter++
		fmt.Println("Pendaftaran berhasil")
	}
}

func cekSudahDaftarDokter(d DataDokter) bool {
	for i := 0; i < jumlahDokter; i++ {
		if dokterBaru[i].Email == d.Email {
			return true
		}
	}
	return false
}

func loginPasien() {
	var loginEmail, loginPassword string

	fmt.Println("Masukkan email anda :")
	fmt.Scan(&loginEmail)

	fmt.Println("Masukkan password anda :")
	fmt.Scan(&loginPassword)

	if cekLoginPasien(loginEmail, loginPassword) {
		fmt.Println("Login berhasil")
		forumKonsultasi()
	} else {
		fmt.Println("Login gagal")
	}
}

func cekLoginPasien(loginEmail, loginPassword string) bool {
	for i := 0; i < jumlahPasien; i++ {
		if pasienBaru[i].Email == loginEmail && pasienBaru[i].Password == loginPassword {
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
		forumKonsultasi()
	} else {
		fmt.Println("Login gagal")
	}
}

func cekLoginDokter(loginEmail, loginPassword string) bool {
	for i := 0; i < jumlahDokter; i++ {
		if dokterBaru[i].Email == loginEmail && dokterBaru[i].Password == loginPassword {
			return true
		}
	}
	return false
}

func menuKonsultasi(pilihKonsultasi *int) {
	fmt.Println("-----------------------")
	fmt.Println("*   MENU KONSULTASI   *")
	fmt.Println("-----------------------")
	fmt.Println("1. Lihat pertanyaan")
	fmt.Println("2. Konsultasi dengan dokter pilihan")
	fmt.Println("3. Kembali ke menu utama")

	fmt.Println("Silahkan pilih menu : 1/2/3")
	fmt.Scan(pilihKonsultasi)
}

func forumKonsultasi() {
	var pilihKonsultasi int

	for {
		menuKonsultasi(&pilihKonsultasi)

		switch pilihKonsultasi {
		case 1:
			lihatPertanyaan()
		case 2:
			konsultasiDokterPilihan()
		case 3:
			return // Kembali ke menu utama
		default:
			fmt.Println("Pilihan tidak tersedia")
		}
	}
}

func konsultasiDokterPilihan() {
	var pilihDokter int

	fmt.Println("--------------------------------")
	fmt.Println("1. Cari berdasarkan spesialisasi")
	fmt.Println("2. Cari berdasarkan nama dokter")
	fmt.Println("3. Tampilkan semua daftar dokter")
	fmt.Println("4. Kembali ke menu konsultasi")
	fmt.Println("--------------------------------")
	fmt.Println("Silahkan pilih 1/2/3/4 : ")
	fmt.Scan(&pilihDokter)

	switch pilihDokter {
	case 1:
		cariSpesialisasi()
	case 2:
		cariDokter()
	case 3:
		tampilkanSemuaDokter()
	case 4:
		return // Kembali ke menu konsultasi
	default:
		fmt.Println("Pilihan tidak tersedia")
	}
}

func cariSpesialisasi() {
	var spesialisasi, pilih, dokter string

	fmt.Print("Cari Spesialisasi : ")
	fmt.Scan(&spesialisasi)

	fmt.Println("-----------------------")
	fmt.Println("*    Daftar Dokter    *")
	fmt.Println("-----------------------")
	for i := 0; i < jumlahDokter; i++ {
		if dokterBaru[i].Spesialisasi == spesialisasi {
			fmt.Println("-----------------------")
			fmt.Print("Nama Dokter : ")
			fmt.Println(dokterBaru[i].NamaLengkap)
			fmt.Println("-----------------------")
			fmt.Println()
		}
	}

	fmt.Println("Ingin kirim pertanyaan dengan dokter pilihan anda? Y/N ")
	fmt.Scan(&pilih)

	if pilih == "Y" || pilih == "y" {
		fmt.Println("Masukkan nama dokter pilihan anda : ")
		fmt.Scan(&dokter)
		postPertanyaan(dokter)
	} else {
		return // Kembali ke menu konsultasi
	}
}

func cariDokter() {
	var dokter, pilih string
	var tempDokter string // variabel untuk menampung nama dokter

	fmt.Print("Cari Dokter : ")
	fmt.Scan(&dokter)

	for i := 0; i < jumlahDokter; i++ {
		if dokterBaru[i].NamaLengkap == dokter {
			fmt.Println(dokterBaru[i].NamaLengkap)
			tempDokter = dokterBaru[i].NamaLengkap
		}
	}

	fmt.Println("Ingin kirim pertanyaan dengan dokter pilihan anda? Y/N ")
	fmt.Scan(&pilih)

	if pilih == "Y" || pilih == "y" {
		postPertanyaan(tempDokter)
	} else {
		return // Kembali ke menu konsultasi
	}
}

func tampilkanSemuaDokter() {
	var pilih, dokter string

	fmt.Println("-----------------------")
	fmt.Println("*    Daftar Dokter    *")
	fmt.Println("-----------------------")

	for i := 0; i < jumlahDokter; i++ {
		fmt.Println("-----------------------")
		fmt.Print("Nama Dokter : ")
		fmt.Println(dokterBaru[i].NamaLengkap)
		fmt.Print("Spesialisasi : ")
		fmt.Println(dokterBaru[i].Spesialisasi)
		fmt.Println("-----------------------")
		fmt.Println()
	}

	fmt.Println("Ingin kirim pertanyaan dengan dokter pilihan anda? Y/N ")
	fmt.Scan(&pilih)

	if pilih == "Y" || pilih == "y" {
		fmt.Println("Masukkan nama dokter pilihan anda : ")
		fmt.Scan(&dokter)
		postPertanyaan(dokter)
	} else {
		return // Kembali ke menu konsultasi
	}

}

func postPertanyaan(namaDokter string) {
	var tanya DataKonsultasi

	fmt.Println("Masukkan pertanyaan anda : ")
	fmt.Scan(&tanya.Pertanyaan)
	// Implementasi post pertanyaan

	pertanyaanBaru[jumlahPertanyaan] = tanya
	jumlahPertanyaan++

	fmt.Println("Pertanyaan anda berhasil dikirim ke dokter " + namaDokter)

}

func lihatPertanyaan() {
	// Implementasi melihat pertanyaan

	fmt.Println("---------------------------")
	fmt.Println("*    Daftar Pertanyaan    *")
	fmt.Println("---------------------------")
	for i := 0; i < jumlahPertanyaan; i++ {
		fmt.Println("---------------------------")
		fmt.Print("Pertanyaan : ")
		fmt.Println(pertanyaanBaru[i].Pertanyaan)
		fmt.Print("Tanggapan : ")
		fmt.Println(pertanyaanBaru[i].Tanggapan)
		fmt.Println("---------------------------")
		fmt.Println()
	}
}
