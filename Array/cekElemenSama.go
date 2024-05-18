package main

import "fmt"

const NMax int = 1000

type tabString [NMax]string

func main() {
	var T tabString
	var n int
	fmt.Scan(&n)
	masukanArray(&T, n)
	fmt.Println(isUniform(T, n))
}

func masukanArray(T *tabString, n int) {
	/*I.S. Terdefinisi nilai bilangan bulat n. Data string sebanyak n buah tersedia
	  pada input device*/
	/*F.S. Array T berisi data yang diberikan*/
	/*Petunjuk : Lakukan perulangan sebanyak n kali untuk melakukan input terhadap
	  data nilai mahasiswa*/
	var i int
	for i = 0; i < n; i++ {
		fmt.Scan(&T[i])
	}
}

func isUniform(T tabString, n int) bool {
	/*Mengembalikan boolean true jika isi data pada array T adalah seragam, dan
	  false jika tidak*/
	var i int
	for i = 0; i < n; i++ {
		if T[i] != T[0] {
			return false
		}
	}
	return true
}
