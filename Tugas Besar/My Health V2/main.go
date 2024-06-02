package main

import "fmt"

func main() {
	var pilih int

	for {
		cetakHeader()

		fmt.Println("--------------------------------------------------")
		fmt.Println("|                    MENU UTAMA                  |")
		fmt.Println("--------------------------------------------------")
		fmt.Println("1. Sign Up")
		fmt.Println("2. Log In")
		fmt.Println("3. Forum Konsultasi")
		fmt.Println("4. Exit")
		fmt.Println("--------------------------------------------------")
		fmt.Print("Silahkan pilih menu (1/2/3/4): ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1: // daftar - signup
			var profesiSignup string

			fmt.Print("Daftar sebagai pasien atau dokter? ")
			fmt.Scan(&profesiSignup)

			if profesiSignup == "pasien" || profesiSignup == "Pasien" {
				signupPasien()
			} else if profesiSignup == "dokter" || profesiSignup == "Dokter" {
				signupDokter()
			}

		case 2: // masuk - login
			var profesiLogin string

			fmt.Print("Masuk sebagai pasien atau dokter? ")
			fmt.Scan(&profesiLogin)

			if profesiLogin == "pasien" || profesiLogin == "Pasien" {
				loginPasien()
			} else if profesiLogin == "dokter" || profesiLogin == "Dokter" {
				loginDokter()
			}

		case 3: // konsultasi
			forumKonsultasiGuest()

		case 4: // keluar - logout
			fmt.Println("Terima kasih telah menggunakan layanan kami.")
			return // Exit the loop and end the program
		}
	}
}

func cetakHeader() {
	fmt.Println("--------------------------------------------------")
	fmt.Println("|              Welcome to My Health              |")
	fmt.Println("--------------------------------------------------")
}
