package main

import "fmt"

func main() {
	var nilai float64
	fmt.Scan(&nilai)
	fmt.Println(kelulusan(nilai))
}

func kelulusan(nilai float64) string {
	if nilai > 40 {
		return "lulus"
	} else {
		return "tidak lulus"
	}
}
