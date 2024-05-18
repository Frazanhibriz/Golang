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