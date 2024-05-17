package main

import "fmt"

const N int = 1000

func main() {
	var x string
	fmt.Scan(&x)

	fmt.Println(searchVokal(x))
}

func searchVokal(x string) bool {
	for i := 0; i < len(x); i++ {
		if x[i] == 'a' || x[i] == 'i' || x[i] == 'u' || x[i] == 'e' || x[i] == 'o' {
			return true
		}
	}
	return false

}
