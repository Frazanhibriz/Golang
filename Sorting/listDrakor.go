package main

import (
	"fmt"
)

// Tipe bentukan struktur tDrakor dengan atribut:
// judul (string), rating (real), jumlah episode (integer), dan durasi (integer)
type tDrakor struct {
	judul         string
	rating        float64
	jumlahEpisode int
	durasi        int
}

// Konstanta NMAX dengan nilai 5
const NMAX = 5

// Tipe alias tabDrakor untuk array of tDrakor dengan ukuran NMAX
type tabDrakor [NMAX]tDrakor

func main() {
	// Deklarasi variabel drakor bertipe array of tDrakor
	var drakor tabDrakor

	// Deklarasi variabel nDrakor bertipe integer untuk menampung banyaknya data drakor
	var nDrakor int

	// Pemanggilan prosedur bacaDataDrakor
	bacaDataDrakor(&drakor, &nDrakor)

	// Pemanggilan prosedur cetakDataDrakor.
	fmt.Println("Data sebelum diurutkan:")
	cetakDataDrakor(drakor, nDrakor)

	// Pemanggilan prosedur urutMenurun
	urutMenurun(&drakor, nDrakor)

	// Pemanggilan prosedur cetakDataDrakor.
	fmt.Println("Data setelah diurutkan menurun berdasarkan rating:")
	cetakDataDrakor(drakor, nDrakor)

	// Pemanggilan prosedur urutMenaik
	urutMenaik(&drakor, nDrakor)

	// Pemanggilan prosedur cetakDataDrakor.
	fmt.Println("Data setelah diurutkan menaik berdasarkan durasi:")
	cetakDataDrakor(drakor, nDrakor)
}

func bacaDataDrakor(A *tabDrakor, n *int) {
	/*
		IS: Array A sebanyak n terdefinisi sembarang
		Proses: 1) Membaca n. Nilai n tidak boleh melebihi NMAX.
		2) Membaca n data array A dengan atribut-atributnya.
		FS: Array A sebanyak n elemen berisi nilai
	*/

	fmt.Scan(n)
	if *n > NMAX {
		*n = NMAX
	}

	for i := 0; i < *n; i++ {

		fmt.Scan(&A[i].judul)

		fmt.Scan(&A[i].rating)

		fmt.Scan(&A[i].jumlahEpisode)

		fmt.Scan(&A[i].durasi)
	}
}

func cetakDataDrakor(A tabDrakor, n int) {
	/*
		IS: Terdefinisi array A sebanyak n elemen
		FS: Tercetak di layar data array A dengan format
		"
		Judul Rating
		Jum Ep Durasi
		<judul> <rating> <jum episode> <durasi>
		..."
		Gunakan format string dengan kolom 20, 6, 6, dan 6
		untuk masing-masing judul, rating, jumlah episode,
		dan durasi.
	*/
	fmt.Printf("%-20s %-6s %-6s %-6s\n", "Judul", "Rating", "Jum Ep", "Durasi")
	for i := 0; i < n; i++ {
		fmt.Printf("%-20s %-6.2f %-6d %-6d\n", A[i].judul, A[i].rating, A[i].jumlahEpisode, A[i].durasi)
	}
}

func urutMenaik(A *tabDrakor, n int) {
    /*
        IS: Terdefinisi Array A sebanyak n elemen
        FS: Array A sebanyak n elemen terurut menaik berdasarkan
        durasi dengan menggunakan algoritma insertion sort
        dengan menggunakan while loop
    */
    i := 1
    for i < n { // While equivalent
        key := A[i]
        j := i - 1
        for j >= 0 && A[j].durasi > key.durasi { // While equivalent
            A[j+1] = A[j]
            j--
        }
        A[j+1] = key
        i++
    }
}

func urutMenurun(A *tabDrakor, n int) {
    /*
        IS: Terdefinisi Array A sebanyak n elemen
        FS: Array A sebanyak n elemen terurut menurun berdasarkan
        rating dengan menggunakan algoritma selection sort
        dengan menggunakan while loop
    */
    i := 0
    for i < n-1 { // While equivalent
        maxIdx := i
        j := i + 1
        for j < n { // While equivalent
            if A[j].rating > A[maxIdx].rating {
                maxIdx = j
            }
            j++
        }
        A[i], A[maxIdx] = A[maxIdx], A[i]
        i++
    }
}

