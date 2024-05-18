package main

import "fmt"

func main() {
	var i, N, poin int
	var goldUser, platinumUser, silverUser int
	fmt.Print("Jumlah pengguna : ")
	fmt.Scanln(&N)

	for i = 0; i < N; i++ {
		fmt.Print("Masukkan poin pengguna : ")
		fmt.Scanln(&poin)

		for poin < 50 {
			fmt.Print("Poin tidak valid, Masukkan kembali : ")
			fmt.Scanln(&poin)
		}

		if poin >= 50 && poin <= 99 {
			silverUser += 1
		} else if poin >= 100 && poin <= 200 {
			platinumUser += 1
		} else if poin > 200 {
			goldUser += 1
		}
	}

	fmt.Println("Gold Users :", goldUser)
	fmt.Println("Platinum Users :", platinumUser)
	fmt.Println("Silver Users :", silverUser)
}
