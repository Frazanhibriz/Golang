package main

import (
	"fmt"
)

func main() {
	type pemain struct {
		name         string
		goal, assist int
	}

	var p pemain
	var topGoal, topAssist string
	var maxGoal, maxAssist int

	for i := 0; i < 3; i++ {
		fmt.Scan(&p.name, &p.goal, &p.assist)
		if p.goal > maxGoal {
			maxGoal = p.goal
			topGoal = p.name
		}
		if p.assist > maxAssist {
			maxAssist = p.assist
			topAssist = p.name
		}
	}

	fmt.Println(topGoal)
	fmt.Println(topAssist)
}
