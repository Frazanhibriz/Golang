package main

import "fmt"

func main() {
	var n, m, i, fpb, kpk int
	fmt.Scan(&n, &m)
	for i = 1; i <= m; i++ {
		if n%i == 0 && m%i == 0 {
			fpb = i
		}
	}
	kpk = (n * m) / fpb // syarat kpk adalah n di kali dengan m lalu hasilnya dibagi dengan fpb
	fmt.Println(kpk)
}
// Billy Task //