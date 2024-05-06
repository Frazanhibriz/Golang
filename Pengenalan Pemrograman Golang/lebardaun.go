package main

import "fmt"

func main() {
	var n, bil, i, max int
	fmt.Scan(&n, &bil)
	max = bil
	i = 1
	for i < n {
		fmt.Scan(&bil)
		if max < bil {
			max = bil
		}
		i++
	}
	fmt.Println(max)
}
