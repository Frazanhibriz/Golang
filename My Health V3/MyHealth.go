package main

import "fmt"

const NMAX int = 1000

type Pengguna struct {
	Nama           string
	Email          string
	Password       string
	Role           string
	Specialization string
}

type DataKonsultasi struct {
	Identitas       Pengguna
	Pertanyaan      string
	Tanggapan       [NMAX]DataRespon
	JumlahTanggapan int
	Tags            string
}

type DataRespon struct {
	Identitas Pengguna
	Response  string
}

var JumlahUser int
var JumlahKonsultasi int

var LoggedInUser Pengguna

var User [NMAX]Pengguna
var Konsultasi [NMAX]DataKonsultasi

func main() {
	var pilih int

	for {
		fmt.Println("------------------------------------------")
		fmt.Println("|  Selamat Datang di Aplikasi My Health  |")
		fmt.Println("------------------------------------------")
		fmt.Println("|               Menu Utama               |")
		fmt.Println("------------------------------------------")
		fmt.Println("1. Daftar")
		fmt.Println("2. Masuk")
		fmt.Println("3. Lihat Forum Konsultasi")
		fmt.Println("4. Keluar")
		fmt.Println("------------------------------------------")
		fmt.Print("Masukkan Pilihan Anda : ")
		fmt.Scan(&pilih)
		fmt.Println("------------------------------------------")

		switch pilih {
		case 1:
			Daftar()
		case 2:
			Masuk()
		case 3:
			LihatForum()
		case 4:
			return
		default:
			fmt.Println("Pilihan Tidak Tersedia")
		}
	}
}

func Daftar() {
	// I.S : Menerima input data pengguna baru (nama, email, password, role, dan spesialisasi jika role adalah dokter).
	// F.S : Menambahkan pengguna baru ke dalam array User jika email dan password belum terdaftar sebelumnya, serta menampilkan pesan pendaftaran berhasil atau gagal.
	var daftar Pengguna

	fmt.Println("------------------------------------------")
	fmt.Println("|                Daftar                  |")
	fmt.Println("------------------------------------------")
	fmt.Print("Masukkan nama anda : ")
	fmt.Scan(&daftar.Nama)
	fmt.Print("Masukkan email anda : ")
	fmt.Scan(&daftar.Email)
	fmt.Print("Masukkan password anda : ")
	fmt.Scan(&daftar.Password)
	fmt.Print("Apakah anda seorang pasien atau dokter? : ")
	fmt.Scan(&daftar.Role)
	if daftar.Role == "dokter" || daftar.Role == "Dokter" {
		fmt.Print("Masukkan spesialisasi anda : ")
		fmt.Scan(&daftar.Specialization)
	}
	fmt.Println("------------------------------------------")

	if cekSudahDaftar(daftar.Email, daftar.Password) {
		fmt.Println("Email Sudah Terdaftar")
	} else {
		User[JumlahUser] = daftar
		JumlahUser++
		fmt.Println("Pendaftaran Berhasil")
	}
}

func cekSudahDaftar(Email, Password string) bool {
	// I.S : Menerima email dan password pengguna.
	// F.S : Mengembalikan true jika email atau password sudah terdaftar, false jika tidak.
	for i := 0; i < JumlahUser; i++ {
		if User[i].Email == Email || User[i].Password == Password {
			return true
		}
	}
	return false
}

func Masuk() {
	// I.S : Menerima input status (dokter atau pasien), email, dan password pengguna.
	// F.S : Memvalidasi data login pengguna dan menampilkan menu konsultasi sesuai status pengguna (pasien atau dokter) jika berhasil masuk, atau pesan gagal masuk.
	var email, password, status string

	fmt.Println("------------------------------------------")
	fmt.Println("|                 Masuk                  |")
	fmt.Println("------------------------------------------")
	fmt.Print("Masuk sebagai dokter atau pasien? : ")
	fmt.Scan(&status)
	fmt.Print("Masukkan email anda : ")
	fmt.Scan(&email)
	fmt.Print("Masukkan password anda : ")
	fmt.Scan(&password)
	fmt.Println("------------------------------------------")

	if cekSudahDaftarPasien(email, password) {
		fmt.Println("Berhasil Untuk Masuk")
		if status == "pasien" || status == "Pasien" {
			MenuKonsultasiPasien()
		} else if status == "dokter" || status == "Dokter" {
			MenuKonsultasiDokter()
		}
	} else {
		fmt.Println("Gagal Untuk Masuk")
	}
}

func cekSudahDaftarPasien(email, password string) bool {
	// I.S : Menerima email dan password pengguna.
	// F.S : Mengembalikan true jika email dan password cocok dengan data pengguna yang terdaftar, serta mengatur LoggedInUser dengan data pengguna yang sesuai.
	for i := 0; i < JumlahUser; i++ {
		if User[i].Email == email && User[i].Password == password {
			LoggedInUser = User[i]
			return true
		}
	}
	return false
}

func MenuKonsultasiPasien() {
	// I.S : Menampilkan menu konsultasi untuk pasien.
	// F.S : Memanggil fungsi sesuai pilihan pengguna.
	var pilih int

	for {
		fmt.Println("-----------------------------------------------")
		fmt.Println("|               Menu Konsultasi               |")
		fmt.Println("-----------------------------------------------")
		fmt.Println("1. Post pertanyaan")
		fmt.Println("2. Cari pertanyaan")
		fmt.Println("3. Tanggapi pertanyaan")
		fmt.Println("4. Lihat pertanyaan secara acak")
		fmt.Println("5. Lihat pertanyaan secara terurut menaik")
		fmt.Println("6. Lihat pertanyaan secara terurut menurun")
		fmt.Println("7. Kembali")
		fmt.Println("-----------------------------------------------")
		fmt.Print("Masukkan pilihan anda : ")
		fmt.Scan(&pilih)
		fmt.Println("-----------------------------------------------")

		switch pilih {
		case 1:
			PostPertanyaan()
		case 2:
			CariPertanyaan()
		case 3:
			TanggapiPertanyaanPasien()
		case 4:
			LihatPertanyaanAcak()
		case 5:
			LihatPertanyaanMenaik()
		case 6:
			LihatPertanyaanMenurun()
		case 7:
			return
		default:
			fmt.Println("Pilihan tidak tersedia")
		}
	}
}

func MenuKonsultasiDokter() {
	// I.S : Menampilkan menu konsultasi untuk dokter.
	// F.S : Memanggil fungsi sesuai pilihan pengguna.
	var pilih int

	for {
		fmt.Println("-----------------------------------------------")
		fmt.Println("|               Menu Konsultasi               |")
		fmt.Println("-----------------------------------------------")
		fmt.Println("1. Tanggapi pertanyaan")
		fmt.Println("2. Cari pertanyaan")
		fmt.Println("3. Lihat topik terurut berdasarkan jumlah pertanyaannya")
		fmt.Println("4. Lihat pertanyaan secara acak")
		fmt.Println("5. Lihat pertanyaan secara terurut menaik")
		fmt.Println("6. Lihat pertanyaan secara terurut menurun")
		fmt.Println("7. Kembali")
		fmt.Println("-----------------------------------------------")
		fmt.Print("Masukkan pilihan anda : ")
		fmt.Scan(&pilih)
		fmt.Println("-----------------------------------------------")

		switch pilih {
		case 1:
			TanggapiPertanyaan()
		case 2:
			CariPertanyaan()
		case 3:
			LihatTopikTerurut()
		case 4:
			LihatPertanyaanAcak()
		case 5:
			LihatPertanyaanMenaik()
		case 6:
			LihatPertanyaanMenurun()
		case 7:
			return
		default:
			fmt.Println("Pilihan tidak tersedia")
		}
	}
}

func CariPertanyaan() {
	// I.S : Menampilkan menu pencarian pertanyaan.
	// F.S : Memanggil fungsi sesuai pilihan pengguna.
	var pilih int
	var tag string

	fmt.Println("-----------------------------------------------")
	fmt.Println("|             Menu Cari Pertanyaan            |")
	fmt.Println("-----------------------------------------------")
	fmt.Println("1. Cari pertanyaan secara terurut menaik")
	fmt.Println("2. Cari pertanyaan secara terurut menurun")
	fmt.Println("3. Cari pertanyaan secara acak")
	fmt.Println("4. Kembali ke menu konsultasi")
	fmt.Println("-----------------------------------------------")
	fmt.Print("Masukkan pilihan anda : ")
	fmt.Scan(&pilih)
	fmt.Println("-----------------------------------------------")

	switch pilih {
	case 1:
		fmt.Print("Masukkan tag pertanyaan yang ingin dicari : ")
		fmt.Scan(&tag)
		CariPertanyaanMenaik(tag)
	case 2:
		fmt.Print("Masukkan tag pertanyaan yang ingin dicari : ")
		fmt.Scan(&tag)
		CariPertanyaanMenurun(tag)
	case 3:
		fmt.Print("Masukkan tag pertanyaan yang ingin dicari : ")
		fmt.Scan(&tag)
		CariPertanyaanAcak(tag)
	case 4:
		return
	default:
		fmt.Println("Pilihan Tidak Tersedia")
	}
}

func CariPertanyaanMenaik(tag string) {
	// I.S : Memanggil fungsi untuk menampilkan pertanyaan secara terurut menaik berdasarkan tag.
	// F.S : Menampilkan pertanyaan secara terurut menaik berdasarkan tag.
	var tags [NMAX]string
	var tagIndex int

	for i := 0; i < JumlahKonsultasi; i++ {
		tags[tagIndex] = Konsultasi[i].Tags
		tagIndex++
	}

	// Binary Search
	index := binarySearch(tags, tagIndex, tag)
	if index != -1 {
		for i := 0; i < JumlahKonsultasi; i++ {
			if Konsultasi[i].Tags == tag {
				TampilkanPertanyaanDanTanggapan(i)
			}
		}
	} else {
		fmt.Println("Tag tidak ditemukan")
	}
}

func CariPertanyaanMenurun(tag string) {
	// I.S : Memanggil fungsi untuk menampilkan pertanyaan secara terurut menurun berdasarkan tag.
	// F.S : Menampilkan pertanyaan secara terurut menurun berdasarkan tag.
	var tags [NMAX]string
	var tagIndex int

	for i := 0; i < JumlahKonsultasi; i++ {
		tags[tagIndex] = Konsultasi[i].Tags
		tagIndex++
	}

	// Binary Search
	index := binarySearch(tags, tagIndex, tag)
	if index != -1 {
		for i := JumlahKonsultasi - 1; i >= 0; i-- {
			if Konsultasi[i].Tags == tag {
				TampilkanPertanyaanDanTanggapan(i)
			}
		}
	} else {
		fmt.Println("Tag tidak ditemukan")
	}
}

func CariPertanyaanAcak(tag string) {
	// I.S : Memanggil fungsi untuk menampilkan pertanyaan secara acak berdasarkan tag.
	// F.S : Menampilkan pertanyaan secara acak berdasarkan tag.
	var tags [NMAX]string
	var tagIndex int

	for i := 0; i < JumlahKonsultasi; i++ {
		tags[tagIndex] = Konsultasi[i].Tags
		tagIndex++
	}

	// Sequential Search
	index := sequentialSearch(tags, tagIndex, tag)
	if index != -1 {
		for i := 0; i < JumlahKonsultasi; i++ {
			if Konsultasi[i].Tags == tag {
				TampilkanPertanyaanDanTanggapan(i)
			}
		}
	} else {
		fmt.Println("Tag tidak ditemukan")
	}
}

func LihatForum() {
	// I.S : Menampilkan menu untuk melihat forum konsultasi.
	// F.S : Memanggil fungsi sesuai pilihan pengguna (lihat pertanyaan secara acak, terurut menaik, terurut menurun, atau keluar).
	var pilih int

	fmt.Println("---------------------------------------------")
	fmt.Println("|          Daftar Forum Konsultasi          |")
	fmt.Println("---------------------------------------------")
	fmt.Println("1. Lihat pertanyaan secara acak")
	fmt.Println("2. Lihat pertanyaan secara terurut menaik")
	fmt.Println("3. Lihat pertanyaan secara terurut menurun")
	fmt.Println("4. Keluar")
	fmt.Println("---------------------------------------------")
	fmt.Print("Masukkan Pilihan Anda : ")
	fmt.Scan(&pilih)

	switch pilih {
	case 1:
		LihatPertanyaanAcak()
	case 2:
		LihatPertanyaanMenaik()
	case 3:
		LihatPertanyaanMenurun()
	case 4:
		return
	default:
		fmt.Println("Pilihan Tidak Tersedia")
	}
}

func TampilkanPertanyaanDanTanggapan(index int) {
	// I.S : index adalah indeks dari pertanyaan yang akan ditampilkan.
	// F.S : Menampilkan pertanyaan dan tanggapan dari indeks yang diberikan.
	fmt.Printf("%d. Pertanyaan: %s ; Penanya: %s (%s)\n", (index + 1), Konsultasi[index].Pertanyaan, Konsultasi[index].Identitas.Nama, Konsultasi[index].Identitas.Role)
	for j := 0; j < Konsultasi[index].JumlahTanggapan; j++ {
		fmt.Printf("\tTanggapan %d: %s ; Penanggap: %s (%s)\n", (j + 1), Konsultasi[index].Tanggapan[j].Response, Konsultasi[index].Tanggapan[j].Identitas.Nama, Konsultasi[index].Tanggapan[j].Identitas.Role)
	}
}

func LihatPertanyaanAcak() {
	// I.S : Daftar konsultasi tersedia.
	// F.S : Menampilkan semua pertanyaan dan tanggapannya secara acak.
	for i := 0; i < JumlahKonsultasi; i++ {
		TampilkanPertanyaanDanTanggapan(i)
	}
}

func LihatPertanyaanMenaik() {
	// I.S : Daftar konsultasi tersedia.
	// F.S : Menampilkan semua pertanyaan dan tanggapannya dalam urutan menaik berdasarkan tag.
	for i := 0; i < JumlahKonsultasi-1; i++ {
		minIndex := i
		for j := i + 1; j < JumlahKonsultasi; j++ {
			if Konsultasi[j].Tags < Konsultasi[minIndex].Tags {
				minIndex = j
			}
		}
		Konsultasi[i], Konsultasi[minIndex] = Konsultasi[minIndex], Konsultasi[i]
	}

	for i := 0; i < JumlahKonsultasi; i++ {
		TampilkanPertanyaanDanTanggapan(i)
	}
}

func LihatPertanyaanMenurun() {
	// I.S : Daftar konsultasi tersedia.
	// F.S : Menampilkan semua pertanyaan dan tanggapannya dalam urutan menurun berdasarkan tag.
	for i := 1; i < JumlahKonsultasi; i++ {
		key := Konsultasi[i]
		j := i - 1

		for j >= 0 && Konsultasi[j].Tags < key.Tags {
			Konsultasi[j+1] = Konsultasi[j]
			j = j - 1
		}
		Konsultasi[j+1] = key
	}

	for i := 0; i < JumlahKonsultasi; i++ {
		TampilkanPertanyaanDanTanggapan(i)
	}
}

func PostPertanyaan() {
	// I.S : Pengguna ingin memposting pertanyaan baru.
	// F.S : Pertanyaan baru ditambahkan ke dalam daftar konsultasi.
	var question, tag string

	fmt.Println("-----------------------------------------------")
	fmt.Println("|                Post Pertanyaan              |")
	fmt.Println("-----------------------------------------------")
	fmt.Print("Hai, mau tanya apa hari ini? : ")
	fmt.Scan(&question)
	fmt.Println("Berikan tag pertanyaan kamu dengan suatu kata agar mudah untuk pencarian pertanyaanmu!")
	fmt.Print("Masukkan tag pertanyaan : ")
	fmt.Scan(&tag)
	fmt.Println("-----------------------------------------------")
	fmt.Println("Pertanyaan berhasil di kirim")

	Konsultasi[JumlahKonsultasi].Pertanyaan = question
	Konsultasi[JumlahKonsultasi].Tags = tag
	Konsultasi[JumlahKonsultasi].Identitas = LoggedInUser
	Konsultasi[JumlahKonsultasi].JumlahTanggapan = 0
	JumlahKonsultasi++
}

func TanggapiPertanyaan() {
	// I.S : Pengguna ingin menanggapi pertanyaan yang sudah ada.
	// F.S : Tanggapan ditambahkan ke pertanyaan yang dipilih oleh pengguna.
	var pilihPertanyaan int
	var response string

	fmt.Println("-----------------------------------------------")
	fmt.Println("|             Tanggapi Pertanyaan             |")
	fmt.Println("-----------------------------------------------")
	fmt.Println("Daftar Pertanyaan:")

	for i := 0; i < JumlahKonsultasi; i++ {
		fmt.Printf("%d. Pertanyaan: %s ; Penanya: %s (%s)\n", (i + 1), Konsultasi[i].Pertanyaan, Konsultasi[i].Identitas.Nama, Konsultasi[i].Identitas.Role)
	}

	fmt.Print("Pilih nomor pertanyaan yang ingin ditanggapi: ")
	fmt.Scan(&pilihPertanyaan)

	if pilihPertanyaan > 0 && pilihPertanyaan <= JumlahKonsultasi {
		fmt.Print("Masukkan tanggapan anda: ")
		fmt.Scan(&response)

		k := &Konsultasi[pilihPertanyaan-1]
		k.Tanggapan[k.JumlahTanggapan].Identitas = LoggedInUser
		k.Tanggapan[k.JumlahTanggapan].Response = response
		k.JumlahTanggapan++

		fmt.Println("Tanggapan berhasil dikirim")
	} else {
		fmt.Println("Pilihan tidak valid")
	}
}

func TanggapiPertanyaanPasien() {
	// I.S : Pengguna ingin menanggapi pertanyaan yang sudah ada.
	// F.S : Tanggapan ditambahkan ke pertanyaan yang dipilih oleh pengguna.
	var pilihPertanyaan int
	var response string

	fmt.Println("-----------------------------------------------")
	fmt.Println("|             Tanggapi Pertanyaan             |")
	fmt.Println("-----------------------------------------------")
	fmt.Println("Daftar Pertanyaan:")

	for i := 0; i < JumlahKonsultasi; i++ {
		fmt.Printf("%d. Pertanyaan: %s ; Penanya: %s (%s)\n", (i + 1), Konsultasi[i].Pertanyaan, Konsultasi[i].Identitas.Nama, Konsultasi[i].Identitas.Role)
	}

	fmt.Print("Pilih nomor pertanyaan yang ingin ditanggapi: ")
	fmt.Scan(&pilihPertanyaan)

	if pilihPertanyaan > 0 && pilihPertanyaan <= JumlahKonsultasi {
		fmt.Print("Masukkan tanggapan anda: ")
		fmt.Scan(&response)

		k := &Konsultasi[pilihPertanyaan-1]
		k.Tanggapan[k.JumlahTanggapan].Identitas = LoggedInUser
		k.Tanggapan[k.JumlahTanggapan].Response = response
		k.JumlahTanggapan++

		fmt.Println("Tanggapan berhasil dikirim")
	} else {
		fmt.Println("Pilihan tidak valid")
	}
}

func LihatTopikTerurut() {
	// I.S : Daftar konsultasi tersedia.
	// F.S : Menampilkan topik terurut berdasarkan jumlah pertanyaannya.

	fmt.Println("-----------------------------------------------")
	fmt.Println("|          Topik Terurut Berdasarkan          |")
	fmt.Println("|           Jumlah Pertanyaannya              |")
	fmt.Println("-----------------------------------------------")

	for i := 0; i < JumlahKonsultasi; i++ {
		maxIndex := i
		for j := i + 1; j < JumlahKonsultasi; j++ {
			if jumlahPertanyaan(Konsultasi[j].Tags) > jumlahPertanyaan(Konsultasi[maxIndex].Tags) {
				maxIndex = j
			}
		}
		// Tampilkan tag dan jumlah pertanyaan
		fmt.Printf("Tag: %s ; Jumlah Pertanyaan: %d\n", Konsultasi[maxIndex].Tags, jumlahPertanyaan(Konsultasi[maxIndex].Tags))
		// Tukar posisi elemen
		Konsultasi[i], Konsultasi[maxIndex] = Konsultasi[maxIndex], Konsultasi[i]
	}
}

func jumlahPertanyaan(tag string) int {
	// Mengembalikan jumlah pertanyaan dengan tag yang diberikan
	count := 0
	for i := 0; i < JumlahKonsultasi; i++ {
		if Konsultasi[i].Tags == tag {
			count++
		}
	}
	return count
}

func binarySearch(array [NMAX]string, size int, target string) int {
	// I.S : array adalah array yang akan dicari, size adalah ukuran array, target adalah elemen yang dicari.
	// F.S : Mengembalikan indeks elemen target dalam array jika ditemukan, -1 jika tidak ditemukan.
	left := 0
	right := size - 1

	for left <= right {
		mid := (left + right) / 2
		if array[mid] == target {
			return mid
		} else if array[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func sequentialSearch(array [NMAX]string, size int, target string) int {
	// I.S : array adalah array yang akan dicari, size adalah ukuran array, target adalah elemen yang dicari.
	// F.S : Mengembalikan indeks elemen target dalam array jika ditemukan, -1 jika tidak ditemukan
	for i := 0; i < size; i++ {
		if array[i] == target {
			return i
		}
	}
	return -1
}
