package main

import "fmt"

const NMAX int = 1000

type dataPasien struct {
	namaLengkap, email, password, pertanyaan string
	lahir                                    dataLahir
}

type dataLahir struct {
	tanggal, bulan, tahun int
	kota                  string
}

type dataDokter struct {
	namaLengkap, email, password string
	lahir                        dataLahir
}


type daftarPasien [NMAX]dataPasien
type daftarDokter [NMAX]dataDokter
type manyQuestion [NMAX]string

func main() {
	var pilih int
	var logEmail, logPassword, profesiSignup, profesiLogin string
	var pasien daftarPasien
	var dokter daftarDokter
	var Question manyQuestion
	var countPasien int // Untuk melacak jumlah pengguna pasien yang terdaftar
	var countDokter int // Untuk melacak jumlah pengguna dokter yang terdaftar

	fmt.Println("------------------------------")
	fmt.Println("*    Welcome to My Health    *")
	fmt.Println("------------------------------")

	for {
		menu()
		fmt.Println("Silahkan pilih menu : 1/2/3/4")
		fmt.Scan(&pilih)
		if pilih == 1 {
			fmt.Println("Daftar sebagai pasien atau dokter? (hint : gunakan huruf kapital pada awal kata)")
			fmt.Scan(&profesiSignup)

			if profesiSignup == "Pasien" {
				if countPasien < NMAX {
					signupPasien(&pasien, &countPasien)
				} else {
					fmt.Println("Batas maksimum pengguna telah tercapai.")
				}
			} else if profesiSignup == "Dokter" {
				if countPasien < NMAX {
					signupDokter(&dokter, &countDokter)
				} else {
					fmt.Println("Batas maksimum pengguna telah tercapai.")
				}
			}

		} else if pilih == 2 {
			fmt.Println("Masuk sebagai pasien atau dokter? (hint : gunakan huruf kapital pada awal kata)")
			fmt.Scan(&profesiLogin)

			if profesiLogin == "Pasien" {
				loginPasien(&pasien, countPasien, &logEmail, &logPassword)
			} else if profesiLogin == "Dokter" {
				loginDokter(&dokter, countDokter, &logEmail, &logPassword)
			}

		} else if pilih == 3 {
			forumKonsultasi(pasien, Question, countPasien)
		} else if pilih == 4 {
			fmt.Println("Terima kasih telah menggunakan layanan kami.")
			break
		}
	}
}

func forumKonsultasi(p daftarPasien, Q manyQuestion, cp int) {
	var pilihKonsultasi, countQ int

	for {
		menuKonsultasi()
		fmt.Println("Silahkan pilih menu : 1/2/3")
		fmt.Scan(&pilihKonsultasi)

		if pilihKonsultasi == 1 {
			ajukanPertanyaan(&p, &Q, &countQ, cp)
		} else if pilihKonsultasi == 2 {
			lihatPertanyaan(Q, countQ)
		} else if pilihKonsultasi == 3 {
			menu()
			break
		}
	}
}

func ajukanPertanyaan(p *daftarPasien, Q *manyQuestion, countQ *int, cp int) {

	fmt.Println("Silahkan ajukan pertanyaan anda!")
	*countQ++
	fmt.Scan(&p[cp].pertanyaan)
	Q[*countQ] = p[cp].pertanyaan
	fmt.Println("Pertanyaan anda berhasil terkirim")

}

func lihatPertanyaan(Q manyQuestion, countQ int) {

	fmt.Println("Daftar pertanyaan :")
	for i := 1; i <= countQ; i++ {
		fmt.Println(Q[i])
	}
}

func menu() {
	fmt.Println("------------------")
	fmt.Println("*   MENU UTAMA   *")
	fmt.Println("------------------")
	fmt.Println("1. Sign Up")
	fmt.Println("2. Log In")
	fmt.Println("3. Forum Konsultasi")
	fmt.Println("4. Exit")
	fmt.Println("------------------")
}

func signupPasien(p *daftarPasien, cp *int) {
	fmt.Println("Masukkan nama lengkap : ")
	fmt.Scan(&p[*cp].namaLengkap)
	fmt.Println("Masukkan email : ")
	fmt.Scan(&p[*cp].email)
	fmt.Println("Masukkan password : ")
	fmt.Scan(&p[*cp].password)
	fmt.Println("Masukkan tanggal lahir : ")
	fmt.Scan(&p[*cp].lahir.tanggal)
	fmt.Println("Masukkan bulan lahir : ")
	fmt.Scan(&p[*cp].lahir.bulan)
	fmt.Println("Masukkan tahun lahir : ")
	fmt.Scan(&p[*cp].lahir.tahun)
	fmt.Println("Masukkan kota lahir : ")
	fmt.Scan(&p[*cp].lahir.kota)

	*cp++
}

func signupDokter(d *daftarDokter, cd *int) {
	fmt.Println("Masukkan nama lengkap : ")
	fmt.Scan(&d[*cd].namaLengkap)
	fmt.Println("Masukkan email : ")
	fmt.Scan(&d[*cd].email)
	fmt.Println("Masukkan password : ")
	fmt.Scan(&d[*cd].password)
	fmt.Println("Masukkan tanggal lahir : ")
	fmt.Scan(&d[*cd].lahir.tanggal)
	fmt.Println("Masukkan bulan lahir : ")
	fmt.Scan(&d[*cd].lahir.bulan)
	fmt.Println("Masukkan tahun lahir : ")
	fmt.Scan(&d[*cd].lahir.tahun)
	fmt.Println("Masukkan kota lahir : ")
	fmt.Scan(&d[*cd].lahir.kota)

	*cd++
}

func loginPasien(p *daftarPasien, cp int, logEmail, logPassword *string) {
	fmt.Println("Masukkan email anda :")
	fmt.Scan(logEmail)
	fmt.Println("Masukkan password anda :")
	fmt.Scan(logPassword)

	if cekSudahDaftarPasien(p, cp, *logEmail, *logPassword) {
		fmt.Println("Login berhasil")
	} else {
		fmt.Println("Login gagal")
	}
}

func loginDokter(d *daftarDokter, cd int, logEmail, logPassword *string) {
	fmt.Println("Masukkan email anda :")
	fmt.Scan(logEmail)
	fmt.Println("Masukkan password anda :")
	fmt.Scan(logPassword)

	if cekSudahDaftarDokter(d, cd, *logEmail, *logPassword) {
		fmt.Println("Login berhasil")
	} else {
		fmt.Println("Login gagal")
	}
}

func cekSudahDaftarPasien(p *daftarPasien, cp int, logEmail, logPassword string) bool {
	for i := 0; i < cp; i++ {
		if p[i].email == logEmail {
			if p[i].password == logPassword {
				return true
			} else {
				fmt.Println("Password salah")
				return false
			}
		}
	}
	fmt.Println("Email tidak terdaftar")
	return false
}

func cekSudahDaftarDokter(d *daftarDokter, cd int, logEmail, logPassword string) bool {
	for i := 0; i < cd; i++ {
		if d[i].email == logEmail {
			if d[i].password == logPassword {
				return true
			} else {
				fmt.Println("Password salah")
				return false
			}
		}
	}
	fmt.Println("Email tidak terdaftar")
	return false
}

func menuKonsultasi() {
	fmt.Println("-----------------------")
	fmt.Println("*   MENU KONSULTASI   *")
	fmt.Println("-----------------------")
	fmt.Println("1. Ajukan pertanyaan")
	fmt.Println("2. Lihat pertanyaan")
	fmt.Println("3. Kembali ke menu utama")
}
