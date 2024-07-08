package main

import "fmt"

func slice() {
	// slice = potongan data dari array
	// slice = data yang memiliki tipe data yang sama

	// slice
	fruits := []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits)

	// akses data slice
	fmt.Println(fruits[0])

	// len() => untuk mengetahui panjang slice
	fmt.Println(len(fruits))

	// append() => menambahkan data ke dalam slice
	// append(slice, data)
	fruits = append(fruits, "papaya")

	// copy() => menyalin data dari slice
	newFruits := make([]string, 2)
	copy(newFruits, fruits) // hanya mengambil 2 data dari fruits
	fmt.Println(newFruits)

	// slice operator
	aFruits := fruits[0:2]
	fmt.Println(aFruits) // ["apple", "grape"]

	// slice operator
	bFruits := fruits[1:]
	fmt.Println(bFruits) // ["grape", "banana", "melon", "papaya"]

	// slice operator
	cFruits := fruits[:3]
	fmt.Println(cFruits) // ["apple", "grape", "banana"]

	// slice operator
	dFruits := fruits[:]
	fmt.Println(dFruits) // ["apple", "grape", "banana", "melon", "papaya"]

	// delete data slice
	// variadic adalah parameter yang bisa menerima lebih dari satu input
	fruits = append(fruits[:2], fruits[3:]...) // ... merupakan variadic yang berfungsi untuk mengeluarkan data dari slice lalu memasukkan data yang lain
	fmt.Println(fruits) // ["apple", "grape", "melon", "papaya"]

	// update data slice
	fruits[0] = "pineapple"
	fmt.Println(fruits) // ["pineapple", "grape", "melon",

	// slice multidimensi
	numbers := [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(numbers)

	// make() => membuat slice baru
	// make(tipe data, panjang, kapasitas)
	// kapasitas = panjang maksimal slice
	newFruits = make([]string, 2, 3)
	fmt.Println(newFruits) 


}
