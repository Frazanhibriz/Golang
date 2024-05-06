package main

import "fmt"

func main() {
	var n, m, i int
	fmt.Scan(&n, &m)
	for i = 1; i <= m; i++ { //karena faktor persekutuan tidak akan pernah lebih besar dari n dan m, maka batas iterasinya bisa n maupun m //
		if n%i == 0 && m%i == 0 {
			fmt.Print(i, " ")
		}
	}
}

// billy task //