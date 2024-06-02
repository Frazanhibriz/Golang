package main

import "fmt"

func signUpForm() {
	var role string

	formTitle("REGISTRASI PENGGUNA")
	formInput("Sebagai pasien atau dokter (p/d)", &role)

	if role == "pasien" || role == "p" {
		if registeredPatient == MAX_PATIENT {
			fmt.Println("Pendaftaran tidak dapat dilanjutkan. Penuh")
			return
		}
		p := User{
			Role: "Pasien",
		}
		formInput("Nama", &p.Name)
		formInput("Umur", &p.Age)
		formInput("Email", &p.Email)
		formInput("Password", &p.Password)
		PatientList[registeredPatient] = p
		registeredPatient++
	} else if role == "dokter" || role == "d" {
		if registeredDoctor == MAX_DOCTOR {
			fmt.Println("Pendaftaran tidak dapat dilanjutkan. Penuh")
			return
		}
		d := User{
			Role: "Dokter",
		}
		formInput("Nama", &d.Name)
		formInput("Umur", &d.Age)
		formInput("Email", &d.Email)
		formInput("Password", &d.Password)
		formInput("Spesialisasi", &d.Specialization)
		DoctorList[registeredDoctor] = d
		registeredDoctor++
	}
	fmt.Println("Selamat, Anda sudah berhasil diregistrasi. Silahkan login di menu sebelumnya")
}

func loginForm() {
	var email string
	var password string

	formTitle("LOGIN")
	formInput("Email", &email)
	formInput("Password", &password)

	for i := 0; i < registeredPatient; i++ {
		if PatientList[i].Email == email && PatientList[i].Password == password {
			fmt.Println("Selamat, Anda berhasil login")
			loggedInUser = PatientList[i]
			showMenu("patient")
			return
		}
	}

	for i := 0; i < registeredDoctor; i++ {
		if DoctorList[i].Email == email && DoctorList[i].Password == password {
			fmt.Println("Selamat, Anda berhasil login")
			loggedInUser = DoctorList[i]
			showMenu("doctor")
			return
		}
	}

	fmt.Println("Email atau password salah")
}
