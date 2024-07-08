package main

import "fmt"

func main() {

	angka := 0

	increment := func() {
		fmt.Println("Increment")
		angka++
	}

	increment()
	increment()
	
	fmt.Println(angka)
}