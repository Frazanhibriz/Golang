package main

import "fmt"

func main() {
	var kendaraan string
	var tipeA, tipeB, tipeC int
	for {
		fmt.Scan(&kendaraan)
		if kendaraan != "A" && kendaraan != "B" && kendaraan != "C" {
			break
		}
		if kendaraan == "A" {
			tipeA++
		} else if kendaraan == "B" {
			tipeB++
		} else if kendaraan == "C" {
			tipeC++
		}
	}
	fmt.Println("Tipe A :", tipeA)
	fmt.Println("Tipe B :", tipeB)
	fmt.Println("Tipe C :", tipeC)
}
