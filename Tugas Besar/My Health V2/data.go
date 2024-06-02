package main

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

type DataTanggapan struct {
	Tanggapan string
	Dokter    *DataDokter
	Pasien    *DataPasien
}

type DataKonsultasi struct {
	Pertanyaan string
	Pasien     *DataPasien
	Tanggapan  [NMAX]DataTanggapan
	Dokter     *DataDokter
	Tags       [NMAX]string
}

var jumlahPasien int
var jumlahDokter int
var jumlahPertanyaan int

var pasienSedangLogin *DataPasien
var dokterSedangLogin *DataDokter

//var pasienBaru DaftarPasien
//var dokterBaru DaftarDokter
//var pertanyaanBaru DaftarPertanyaan

var DaftarPasien [NMAX]*DataPasien
var DaftarDokter [NMAX]*DataDokter
var DaftarPertanyaan [NMAX]*DataKonsultasi
