package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	fmt.Println(totalTabungan(x, y, z))
}

func totalTabungan(x, y, z int) int {
	var total, nabung int

	nabung = x
	for i := 1; i <= z; i++ {
		if i%2 == 0 || i%3 == 0 {
			
			total += nabung 
			nabung += y
		}
	}
	return total
}
