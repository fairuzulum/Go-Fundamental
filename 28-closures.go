package main

import "fmt"

func closure() {

	angka := 0

	increment := func() {
		fmt.Println("Increment")
		angka++
	}

	increment()
	increment()
	
	fmt.Println(angka)
}