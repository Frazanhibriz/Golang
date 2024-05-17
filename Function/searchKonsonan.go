package main

import "fmt"

func main() {
	var x string
	fmt.Scan(&x)

	fmt.Println(searchKonsonan(x))
}

func searchKonsonan(x string) bool {
	for i := 0; i < len(x); i++ {
		if x[i] != 'a' && x[i] != 'i' && x[i] != 'u' && x[i] != 'e' && x[i] != 'o' {
			return true
		}
	}
	return false
}
