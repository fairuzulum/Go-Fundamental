package main

import "fmt"

func recursiveFunction() {
	// Recursive Function
	// Recursive function adalah function yang memanggil dirinya sendiri

	// Contoh Recursive Function
	// Faktorial => 5! = 5 * 4 * 3 * 2 * 1

	// Faktorial dengan Loop
	var result = 1
	for i := 5; i > 0; i-- {
		result *= i
	}
	fmt.Println(result)

	// Faktorial dengan Recursive Function
	fmt.Println(faktorial(5))
}

func faktorial(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * faktorial(value-1)
	}
}