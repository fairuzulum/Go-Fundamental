package main

import "fmt"

func typeDataArray() {
	// array = kumpulan data yang memiliki tipe data yang sama

	// deklarasi array
	var names [3]string
	names[0] = "Ahmad"
	names[1] = "Fairuz"
	names[2] = "Ulumuddin"
	fmt.Println(names)
	// akses data array
	fmt.Println(names[1])

	// deklarasi array dengan inisialisasi
	fruits := [4]string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits)

	// array multidimensi
	numbers := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(numbers)

	// len() => untuk mengetahui panjang array
	fmt.Println(len(fruits))

}
