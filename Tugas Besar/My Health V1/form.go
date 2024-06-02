package main

import (
	"fmt"
)

func formTitle(title string) {
	fmt.Printf("\n\n\n")
	fmt.Printf("====================================================\n")
	fmt.Printf(" %s \n", title)
	fmt.Printf("====================================================\n")
}

func formInput(caption string, a ...any) {
	fmt.Printf("%s : ", caption)
	fmt.Scan(a...)
}
