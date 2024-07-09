package main

import "fmt"

func variadic() {
	// Variadic Function adalah function yang memungkinkan kita untuk memasukkan banyak argumen sekaligus
	// Variadic Function ditandai dengan adanya titik tiga (...) sebelum tipe data parameter terakhir
	// Variadic Function hanya bisa memiliki satu variadic parameter
	// Variadic Parameter adalah parameter yang bisa menerima lebih dari satu argumen
	// Variadic Parameter disimpan dalam bentuk array
	// Variadic Parameter harus ditempatkan di posisi terakhir
	// Variadic Parameter tidak wajib diisi
	// Variadic Parameter bisa diisi dengan nol atau lebih argumen
	// Variadic Parameter bisa diisi dengan tipe data apapun, termasuk tipe data yang berbeda-beda

	fmt.Println(sumAll(10, 20, 30, 40, 50)) 

	// memasukan slice ke variadic function
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println(sumAll(numbers...))
}

// Contoh Variadic Function
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
