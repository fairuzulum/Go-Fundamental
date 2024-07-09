package main

import (
	"fmt"
	"golang-dasar/helper"
)

func main() {
	helper.SayHiii("fairuz")

	// access modifier digolang hanya dibedakan dari penulisan awal
	// jika diawali huruf kapital maka bisa diakses dari package lain
	// jik diawali huruf kecil maka tidak bisa diakases
	fmt.Println(helper.Aplication)
	//fmt.Println(helper.version)  //error
}
