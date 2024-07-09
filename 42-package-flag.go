package main

import (
	"flag"
	"fmt"
)

func main() {
	// package flag adalah package yang digunakan untuk mengambil input dari command line
	// package flag memiliki 2 fungsi utama yaitu flag.String() dan flag.Parse()
	// flag.String() digunakan untuk mendeklarasikan variabel yang akan digunakan untuk menyimpan input dari command line
	// flag.Parse() digunakan untuk mengambil input dari command line dan menyimpannya ke variabel yang sudah dideklarasikan

	// contoh penggunaan package flag
	
	host := flag.String("host", "localhost", "Put your database host")
	username := flag.String("username", "root", "Put your database username")
	password := flag.String("password", "root", "Put your database password")

	flag.Parse()

	fmt.Println(*host, *username, *password)


}