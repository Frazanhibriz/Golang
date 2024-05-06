package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var pilih int
	intro()
	for {
		menu_utama(&pilih)
		switch pilih {
		case 1:
			jumlah_bilangan_asli()
		case 2:
			faktorial()
		case 3:
			fibonacci()
		default:
			clear_screen()
		}
		if pilih == 4 {
			break
		}
	}
	bye()
}

func intro() {
	clear_screen()
	fmt.Println("Selamat Datang")
}

func bye() {
	clear_screen()
	fmt.Println("Sampai Jumpa")
}

func clear_screen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func menu_utama(p *int) {
	fmt.Println("---------------------------")
	fmt.Println("           MENU            ")
	fmt.Println("---------------------------")
	fmt.Println("1. Jumlah Bilangan Asli")
	fmt.Println("2. Faktorial")
	fmt.Println("3. Fibonacci")
	fmt.Println("4. Exit")
	fmt.Println("---------------------------")
	fmt.Print("Pilih (1/2/3/4): ")
	fmt.Scan(p)
}

func menu_jumlah_bilangan_asli() {
	clear_screen()
	fmt.Println("---------------------------")
	fmt.Println("           MENU            ")
	fmt.Println("---------------------------")
	fmt.Println("1. Jumlah Bilangan Asli")
	fmt.Println("2. Exit")
	fmt.Println("---------------------------")
}

func jumlah_bilangan_asli() {
	var pilih, N, jumlah int
	var lagi string
	for {
		menu_jumlah_bilangan_asli()
		fmt.Print("Pilih (1/2): ")
		fmt.Scan(&pilih)
		if pilih == 1 {
			for {
				fmt.Print("Masukkan Bilangan Asli N: ")
				fmt.Scan(&N)
				jumlah = jumlah_bilangan_asli_rekursif(N)
				fmt.Printf("Jumlah %d Bilangan Asli Pertama: %d\n", N, jumlah)
				fmt.Print("Lagi (Y/T): ")
				fmt.Scan(&lagi)
				if lagi == "Y" {
					break
				}
			}
			clear_screen()
		} else if pilih == 2 {
			clear_screen()
			break
		}
	}
}

func jumlah_bilangan_asli_rekursif(bil int) int {
	if bil == 0 {
		return 0
	} else {
		return bil + jumlah_bilangan_asli_rekursif(bil-1)
	}
}

func menu_faktorial() {
	clear_screen()
	fmt.Println("---------------------------")
	fmt.Println("           MENU            ")
	fmt.Println("---------------------------")
	fmt.Println("1. Faktorial")
	fmt.Println("2. Exit")
	fmt.Println("---------------------------")
}

func faktorial() {
	var pilih, N, suku int
	var lagi string
	for {
		menu_faktorial()
		fmt.Print("Pilih (1/2): ")
		fmt.Scan(&pilih)
		if pilih == 1 {
			for {
				fmt.Print("Masukkan Bilangan Asli N: ")
				fmt.Scan(&N)
				suku = faktorial_rekursif(N)
				fmt.Printf("Nilai Faktorial %d: %d\n", N, suku)
				fmt.Print("Lagi (Y/T): ")
				fmt.Scan(&lagi)
				if lagi == "Y" {
					break
				}
			}
			clear_screen()
		} else if pilih == 2 {
			clear_screen()
			break
		}
	}
}

func faktorial_rekursif(bil int) int {
	if bil == 1 {
		return 1
	} else {
		return bil * faktorial_rekursif(bil-1)
	}
}

func menu_fibonacci() {
	clear_screen()
	fmt.Println("---------------------------")
	fmt.Println("           MENU            ")
	fmt.Println("---------------------------")
	fmt.Println("1. Fibonacci")
	fmt.Println("2. Exit")
	fmt.Println("---------------------------")
}

func fibonacci() {
	var pilih, N, suku int
	var lagi string
	for {
		menu_fibonacci()
		fmt.Print("Pilih (1/2): ")
		fmt.Scan(&pilih)
		if pilih == 1 {
			for {
				fmt.Print("Masukkan Bilangan Asli N: ")
				fmt.Scan(&N)
				suku = fibonacci_rekursif(N)
				fmt.Printf("Suku Fibonacci ke-%d: %d\n", N, suku)
				fmt.Print("Lagi (Y/T): ")
				fmt.Scan(&lagi)
				if lagi == "Y" {
					break
				}
			}
			clear_screen()
		} else if pilih == 2 {
			clear_screen()
			break
		}
	}
}

func fibonacci_rekursif(bil int) int {
	if bil == 1 {
		return 0
	} else if bil == 2 {
		return 1
	} else if bil == 3 {
		return 1
	} else {
		return fibonacci_rekursif(bil-2) + fibonacci_rekursif(bil-1)
	}
}
