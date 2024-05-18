package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	printDigitsWithDash(n)
}

func printDigitsWithDash(n int) {
	if n < 10 {
		fmt.Print(n)
	} else {
		printDigitsWithDash(n / 10)
		fmt.Print("-")
		fmt.Print(n % 10)
	}
}
