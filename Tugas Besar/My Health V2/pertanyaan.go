package main

import (
	"fmt"
)

func forumKonsultasiPasien() {
	var pilihKonsultasi int

	for {
		fmt.Println("--------------------------------------------------")
		fmt.Println("|                 MENU KONSULTASI                |")
		fmt.Println("--------------------------------------------------")
		fmt.Println("1. Lihat pertanyaan")
		fmt.Println("2. Konsultasi dengan dokter pilihan")
		fmt.Println("3. Log Out")
		fmt.Println("--------------------------------------------------")
		fmt.Print("Silahkan pilih menu: ")
		fmt.Scan(&pilihKonsultasi)

		switch pilihKonsultasi {
		case 1:
			lihatPertanyaan()
		case 2:
			konsultasiDokterPilihan()
		case 3:
			pasienSedangLogin = nil
			return // Kembali ke menu utama
		default:
			fmt.Println("Pilihan tidak tersedia")
		}
	}
}

func forumKonsultasiDokter() {
	var pilihKonsultasi int

	for {
		fmt.Println("--------------------------------------------------")
		fmt.Println("|                 MENU KONSULTASI                |")
		fmt.Println("--------------------------------------------------")
		fmt.Println("1. Lihat pertanyaan")
		fmt.Println("2. Log Out")
		fmt.Println("--------------------------------------------------")
		fmt.Print("Silahkan pilih menu: ")
		fmt.Scan(&pilihKonsultasi)

		switch pilihKonsultasi {
		case 1:
			lihatPertanyaan()
		case 2:
			dokterSedangLogin = nil
			return // Kembali ke menu utama
		default:
			fmt.Println("Pilihan tidak tersedia")
		}
	}
}

func forumKonsultasiGuest() {
	var pilihKonsultasi int

	for {
		fmt.Println("--------------------------------------------------")
		fmt.Println("|                 MENU KONSULTASI                |")
		fmt.Println("--------------------------------------------------")
		fmt.Println("1. Lihat pertanyaan")
		fmt.Println("2. Konsultasi umum")
		fmt.Println("3. Konsultasi dengan dokter pilihan")
		fmt.Println("4. Kembali ke menu utama")
		fmt.Println("--------------------------------------------------")
		fmt.Print("Silahkan pilih menu: ")
		fmt.Scan(&pilihKonsultasi)

		switch pilihKonsultasi {
		case 1:
			lihatPertanyaan()
			break
		case 2:
			postPertanyaan(nil)
			break
		case 3:
			konsultasiDokterPilihan()
			break
		case 4:
			return // Kembali ke menu utama
		default:
			fmt.Println("Pilihan tidak tersedia")
		}
	}
}

func lihatPertanyaan() {
	var pilihMenu int
	var dataKonsul *DataKonsultasi

	for {
		fmt.Println("--------------------------------------------------")
		fmt.Println("|               DAFTAR PERTANYAAN                |")
		fmt.Println("--------------------------------------------------")
		fmt.Println("1. Lihat semua pertanyaan")
		fmt.Println("2. Lihat pertanyaan dengan subyek tertentu")

		if pasienSedangLogin != nil {
			fmt.Println("3. Lihat pertanyaan yang Anda kirim")
		} else if dokterSedangLogin != nil {
			fmt.Println("3. Lihat pertanyaan yang Anda tanggapi")
		}

		fmt.Println("4. Kembali ke menu utama")
		fmt.Println("--------------------------------------------------")
		fmt.Print("Silahkan pilih menu: ")
		fmt.Scan(&pilihMenu)

		switch pilihMenu {
		case 1:
			dataKonsul = cetakSemuaPertanyaan()
			break
		case 2:
			dataKonsul = cetakPertanyaanBerdasarkanTag()
			break
		case 3:
			if pasienSedangLogin != nil {
				dataKonsul = cetakPertanyaanPasien()
			} else if dokterSedangLogin != nil {
				dataKonsul = cetakPertanyaanDokter()
			}
			break
		case 4:
			return

		default:

		}

		if dataKonsul != nil {
			menuKonsultasiDanTanggapan(dataKonsul)
		}
	}
}

func cetakSemuaPertanyaan() *DataKonsultasi {
	var pilih, totalPertanyaan int

	for i := 0; i < jumlahPertanyaan; i++ {
		if DaftarPertanyaan[i].Pertanyaan == "" {
			i = jumlahPertanyaan
		} else {
			fmt.Printf("%d. %s , penanya: %s\n", (i + 1), DaftarPertanyaan[i].Pertanyaan, DaftarPertanyaan[i].Pasien.NamaLengkap)
			totalPertanyaan++
		}
	}

	fmt.Println("--------------------------------------------------")
	if totalPertanyaan == 0 {
		fmt.Print("Maaf, saat ini belum ada pertanyaan yang terdaftar. Tekan enter untuk ke menu sebelumnya")
		fmt.Scan(&pilih)
		return nil
	} else {
		fmt.Print("Silahkan pilih nomor pertanyaan: ")
		fmt.Scan(&pilih)
		return DaftarPertanyaan[pilih-1]
	}
}

func cetakPertanyaanBerdasarkanTag() *DataKonsultasi {
	var pilih int
	var tag1, tag2, tag3 string
	var daftarKonsul [NMAX]*DataKonsultasi
	var idx int

	fmt.Print("Masukkan subyek yang ingin dicari (max. 3 subyek): ")
	fmt.Scanf("%s %s %s", &tag1, &tag2, &tag3)

	for i := 0; i < jumlahPertanyaan; i++ {
		tagFound := false

		//find tags
		for j := 0; j < NMAX; j++ {
			if DaftarPertanyaan[i].Tags[j] == "" {
				j = NMAX

			} else {

				if DaftarPertanyaan[i].Tags[j] == tag1 {
					tagFound = true
				} else if DaftarPertanyaan[i].Tags[j] == tag2 {
					tagFound = true
				} else if DaftarPertanyaan[i].Tags[j] == tag3 {
					tagFound = true
				}
			}
		}

		if tagFound {
			daftarKonsul[idx] = DaftarPertanyaan[i]
			idx++

			fmt.Printf("%d. %s , penanya: %s\n", idx, DaftarPertanyaan[i].Pertanyaan, DaftarPertanyaan[i].Pasien.NamaLengkap)
		}
	}

	fmt.Println("--------------------------------------------------")
	fmt.Print("Silahkan pilih nomor pertanyaan: ")
	fmt.Scan(&pilih)

	return daftarKonsul[pilih-1]
}

func cetakPertanyaanPasien() *DataKonsultasi {
	var daftarKonsul [NMAX]*DataKonsultasi
	var idx, pilih int

	for i := 0; i < jumlahPertanyaan; i++ {
		if DaftarPertanyaan[i].Pasien == pasienSedangLogin {
			daftarKonsul[idx] = DaftarPertanyaan[i]
			idx++

			fmt.Printf("%d. %s , penanya: %s\n", idx, DaftarPertanyaan[i].Pertanyaan, DaftarPertanyaan[i].Pasien.NamaLengkap)
		}
	}

	fmt.Println("-----------------------------------")
	fmt.Scanf("Silahkan pilih nomor pertanyaan: %d", &pilih)

	return daftarKonsul[pilih-1]
}

func cetakPertanyaanDokter() *DataKonsultasi {
	var daftarKonsul [NMAX]*DataKonsultasi
	var idx, pilih int

	for i := 0; i < jumlahPertanyaan; i++ {
		if DaftarPertanyaan[i].Dokter == dokterSedangLogin {
			daftarKonsul[idx] = DaftarPertanyaan[i]
			idx++

			fmt.Printf("%d. %s , penanya: %s\n", idx, DaftarPertanyaan[i].Pertanyaan, DaftarPertanyaan[i].Pasien.NamaLengkap)
		}
	}

	fmt.Println("-----------------------------------")
	fmt.Scanf("Silahkan pilih nomor pertanyaan: %d", &pilih)

	return daftarKonsul[pilih-1]
}

func menuKonsultasiDanTanggapan(dataKonsul *DataKonsultasi) {
	var tanggapanBaru DataTanggapan
	var jumlahTanggapan int

	fmt.Printf("Pertanyaan:\n %s\n", dataKonsul.Pertanyaan)
	fmt.Printf("Nama Pasien: %s\n", dataKonsul.Pasien.NamaLengkap)
	fmt.Printf("Umur Pasien: %d\n", 2024-dataKonsul.Pasien.Lahir.Tahun)
	fmt.Println()

	for i := 0; i < NMAX; i++ {
		if dataKonsul.Tanggapan[i].Tanggapan == "" {
			i = NMAX
		} else {
			jumlahTanggapan++
			if dataKonsul.Tanggapan[i].Dokter != nil {
				fmt.Printf("Dokter %s: %s", dataKonsul.Tanggapan[i].Dokter.NamaLengkap, dataKonsul.Tanggapan[i].Tanggapan)
			} else {
				fmt.Printf("Pasien %s: %s", dataKonsul.Tanggapan[i].Pasien.NamaLengkap, dataKonsul.Tanggapan[i].Tanggapan)
			}
		}
	}

	tanggapanBaru = DataTanggapan{}
	fmt.Println("Silahkan beri tanggapan atau enter untuk kembali ke menu sebelumnya:")
	fmt.Scan(&tanggapanBaru.Tanggapan)

	if tanggapanBaru.Tanggapan != "" {
		if pasienSedangLogin != nil {
			tanggapanBaru.Pasien = pasienSedangLogin
		} else {
			tanggapanBaru.Dokter = dokterSedangLogin
		}
		dataKonsul.Tanggapan[jumlahTanggapan] = tanggapanBaru
	}

	return
}

func postPertanyaan(dokter *DataDokter) {
	if pasienSedangLogin.Email == "" {
		return
	}

	tanya := &DataKonsultasi{
		Dokter: dokter,
		Pasien: pasienSedangLogin,
	}

	fmt.Println("Masukkan pertanyaan Anda: ")
	fmt.Scan(&tanya.Pertanyaan)
	// Implementasi post pertanyaan

	DaftarPertanyaan[jumlahPertanyaan] = tanya
	jumlahPertanyaan++

	if dokter != nil {
		fmt.Println("Pertanyaan anda berhasil dikirim ke dokter " + dokter.NamaLengkap)
	}
}

func konsultasiDokterPilihan() {
	var pilihDokter int

	fmt.Println("--------------------------------------------------")
	fmt.Println("|       KONSULTASI DENGAN DOKTER PILIHAN         |")
	fmt.Println("--------------------------------------------------")
	fmt.Println("1. Cari berdasarkan spesialisasi")
	fmt.Println("2. Cari berdasarkan nama dokter")
	fmt.Println("3. Tampilkan semua daftar dokter")
	fmt.Println("4. Kembali ke menu konsultasi")
	fmt.Println("--------------------------------------------------")
	fmt.Print("Silahkan pilih: ")
	fmt.Scan(&pilihDokter)

	switch pilihDokter {
	case 1:
		cariSpesialisasi()
		break
	case 2:
		cariDokter()
		break
	case 3:
		tampilkanSemuaDokter("", "")
		break
	case 4:
		return // Kembali ke menu konsultasi
	default:
		fmt.Println("Pilihan tidak tersedia")
	}
}
