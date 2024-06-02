package main

import (
	"fmt"
)

func getMenu(menuID string) Menu {
	thisMenu := Menu{}

	switch menuID {
	case "main":
		thisMenu.Title = "MENU UTAMA"
		thisMenu.Choices[0] = ChoiceElement{Caption: "Sign Up", Key: "signup"}
		thisMenu.Choices[1] = ChoiceElement{Caption: "Login", Key: "login"}
		thisMenu.Choices[2] = ChoiceElement{Caption: "Forum Konsultasi", Key: "forum"}
		thisMenu.Choices[3] = ChoiceElement{Caption: "Keluar", Key: "exit"}
	case "forum":
		thisMenu.Title = "FORUM KONSULTASI"
		thisMenu.Choices[0] = ChoiceElement{Caption: "Lihat Daftar Pertanyaan", Key: "questionList"}
		thisMenu.Choices[1] = ChoiceElement{Caption: "Urutkan Pertanyaan Berdasarkan Tag", Key: "sortQuestionByTag"}
		thisMenu.Choices[2] = ChoiceElement{Caption: "Cari Pertanyaan Berdasarkan Tag", Key: "findQuestionByTag"}
		thisMenu.Choices[3] = ChoiceElement{Caption: "Kembali ke Menu Sebelumnya", Key: "back"}
	case "patient":
		thisMenu.Title = "MENU PASIEN"
		thisMenu.Choices[0] = ChoiceElement{Caption: "Lihat Daftar Dokter", Key: "doctorList"}
		thisMenu.Choices[1] = ChoiceElement{Caption: "Cari Dokter Berdasarkan Nama", Key: "findDoctorByName"}
		thisMenu.Choices[2] = ChoiceElement{Caption: "Cari Dokter Berdasarkan Spesialisasi", Key: "findDoctorBySpec"}
		thisMenu.Choices[3] = ChoiceElement{Caption: "Lihat Daftar Pertanyaan", Key: "questionList"}
		thisMenu.Choices[4] = ChoiceElement{Caption: "Urutkan Pertanyaan Berdasarkan Tag", Key: "sortQuestionByTag"}
		thisMenu.Choices[5] = ChoiceElement{Caption: "Cari Pertanyaan Berdasarkan Tag", Key: "findQuestionByTag"}
		thisMenu.Choices[6] = ChoiceElement{Caption: "Kirim Pertanyaan Baru", Key: "postNewQuestion"}
		thisMenu.Choices[7] = ChoiceElement{Caption: "Logout", Key: "logout"}
	case "doctor":
		thisMenu.Title = "MENU DOKTER"
		thisMenu.Choices[0] = ChoiceElement{Caption: "Lihat Daftar Pasien", Key: "patientList"}
		thisMenu.Choices[1] = ChoiceElement{Caption: "Cari Pasien Berdasarkan Nama", Key: "findPatientByName"}
		thisMenu.Choices[2] = ChoiceElement{Caption: "Lihat Daftar Pertanyaan", Key: "questionList"}
		thisMenu.Choices[3] = ChoiceElement{Caption: "Urutkan Pertanyaan Berdasarkan Tag", Key: "sortQuestionByTag"}
		thisMenu.Choices[4] = ChoiceElement{Caption: "Cari Pertanyaan Berdasarkan Tag", Key: "findQuestionByTag"}
		thisMenu.Choices[5] = ChoiceElement{Caption: "Logout", Key: "logout"}
	default:
		//fmt.Println("Invalid Menu ID")
	}

	return thisMenu
}

func showMenu(menuID string) {
	thisMenu := getMenu(menuID)

	keepShowMenu := true
	for keepShowMenu {
		key := printMenu(thisMenu)

		//ßfmt.Pritln(key)

		switch key {
		case "login":
			loginForm()
			break

		case "logout":
			loggedInUser = User{}
			keepShowMenu = false
			break

		case "forum":

		case "signup":
			signUpForm()
			break

		case "postNewQuestion":
			postNewQuestion()
			break

		case "doctorList":
			//doctorList()
			break

		case "patientList":
			//patientList()
			break

		case "questionList":
			questionList()
			break

		case "sortQuestionByTag":
			sortQuestionByTag("asc")
			break

		case "findQuestionByTag":
			//findQuestionByTag()
			break

		case "findDoctorByName":
			//findDoctorByName()
			break

		case "findDoctorBySpec":
			//findDoctorBySpec()
			break

		case "back":
			keepShowMenu = false
		case "exit":
			keepShowMenu = false
		default:
			showMenu(key)
		}
	}
}

func printMenu(thisMenu Menu) string {
	var pilih int

	fmt.Printf("\n\n\n")
	printBorder("=")
	printMenuTitle(thisMenu.Title)
	printBorder("-")

	for i := 0; i < MAX_CHOICE && thisMenu.Choices[i].Caption != ""; i++ {
		fmt.Printf("%d. %s\n", (i + 1), thisMenu.Choices[i].Caption)
	}
	printBorder("-")

	if thisMenu.Prompt == "" {
		thisMenu.Prompt = "Masukkan Pilihan Anda"
	}
	fmt.Printf(" %s: ", thisMenu.Prompt)
	fmt.Scan(&pilih)

	return thisMenu.Choices[pilih-1].Key
}

func printBorder(c string) {
	for i := 0; i < 70; i++ {
		fmt.Print(c)
	}
	fmt.Println("")
}

func printMenuTitle(title string) {
	if loggedInUser.Role == "Dokter" {
		title = fmt.Sprintf("%s   (Pengguna: dr. %s)", title, loggedInUser.Name)
	} else if loggedInUser.Role == "Pasien" {
		title = fmt.Sprintf("%s   (Pengguna: Tn/Ny. %s)", title, loggedInUser.Name)
	}
	lenTitle := len(title)
	for i := 0; i < (70 - 4 - lenTitle); i++ {
		title += " "
	}
	fmt.Printf("| %s |\n", title)
}
