package main

import "fmt"

type DataHarmonis struct {
	JumlahKebalikan float64
	JumlahBilangan  int
}

func rataHarmonis(data DataHarmonis) float64 {
	return float64(data.JumlahBilangan) / data.JumlahKebalikan
}

func main() {
	var data DataHarmonis
	var inputData int
	fmt.Scan(&data.JumlahBilangan)
	for i := 0; i < data.JumlahBilangan; i++ {
	    fmt.Scan(&inputData)
	    data.JumlahKebalikan += 1 / float64(inputData)
	}
	rata := rataHarmonis(data)
	fmt.Printf("%.2f\n", rata)
}
