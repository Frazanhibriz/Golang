package main

import "fmt"

func lowToUpper(kar byte) byte {
	var uppercase byte

	uppercase = kar
	if kar >= 'a' && kar <= 'z' {
		uppercase = kar - 32
	}
	return uppercase
}

func main() {
	var kar byte

	fmt.Scanf("%c", &kar)
	fmt.Printf("%c\n", lowToUpper(kar))
}
