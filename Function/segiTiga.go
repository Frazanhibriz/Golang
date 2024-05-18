package main

import "fmt"

func main() {
	var s1, s2, s3 float64
	fmt.Scan(&s1, &s2, &s3)
	fmt.Println(segitgaSamaKaki(s1, s2, s3))
}

func segitgaSamaKaki(s1, s2, s3 float64) bool {
	if s1 == s2 && s1 != s3 {
		return true
	} else if s1 == s3 && s1 != s2 {
		return true
	} else if s2 == s3 && s2 != s1 {
		return true
	} else {
		return false
	}
}
