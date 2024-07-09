package main

// import "fmt"

type Address struct {
	City, Province, Country string
}

// func main() {
// 	// pointer => adalah reference ke lokasi data di memori
// 	// golang pass by value => artinya saat kita mengirim sebuah data ke sebuah function, sebenarnya yang dikirim adalah duplikasi data tersebut
// 	// jika kita ingin mengubah data asli, kita bisa menggunakan pointer

// 	address := Address{"Bekasi", "Jawa Barat", "Indonesia"}
// 	addressPointer := &address

// 	addressPointer.City = "Jakarta"	

// 	*addressPointer = Address{"Bandung", "Jawa Barat", "Indonesia"}

// 	fmt.Println(addressPointer)
// 	fmt.Println(address)

// 	address2 := new(Address)
// 	address2.City = "Bogor"
// 	address2.Province = "Jawa Barat"
// 	address2.Country = "Indonesia"
// 	fmt.Println(address2)
// }