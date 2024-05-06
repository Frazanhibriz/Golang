package main

import "fmt"

type bowl struct {
	noodle, topping string
	hot             bool
}

func main() {
	var buy bowl

	fmt.Scan(&buy.noodle, &buy.hot)
	hitungHarga(buy)

}
func hitungHarga(buy bowl) {
	price := 0
	if buy.noodle == "kwetiau" {
		price += 8000
	} else if buy.noodle == "bihun" {
		price += 7000
	} else if buy.noodle == "kuning" {
		price += 9000
	}

	if buy.hot == true {
		price += 1000
	}

	for i := 0; i < 5; i++ {
		fmt.Scan(&buy.topping)
		if buy.topping == "ayam" || buy.topping == "pangsit" {
			price += 5000
		} else if buy.topping == "telur" {
			price += 3000
		} else if buy.topping == "sayur" {
			price += 2000
		} else if buy.topping == "pangsit" {
			price += 5000
		} else if buy.topping == "0" {
			break
		}
	}
	fmt.Println(price)
}
