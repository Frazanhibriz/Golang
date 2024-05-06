package main

import "fmt"

func main() {
	var pengunjung, rekapan, data int
	for {
		fmt.Printf("Hari %d : ", data+1)
		fmt.Scan(&pengunjung)
		if pengunjung >= 0 && pengunjung <= 200 {
			data++
			rekapan = rekapan + pengunjung
		}

		if data == 5 {
			break
		}

	}
	fmt.Println(rekapan)
}

// Falah Task //
