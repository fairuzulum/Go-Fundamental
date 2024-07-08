package main

import "fmt"

func breakAndContinue() {
	// break and continue

	// break
	// program akan berhenti ketika i == 3
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}

	// continue
	// program akan skip ketika i == 3
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
}