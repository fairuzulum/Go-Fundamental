package main

import (
	"fmt"
	"strconv"
)

func main() {
	// package strconv adalah package yang digunakan untuk mengkonversi tipe data

	// mengubah string ke bool
	fmt.Println(strconv.ParseBool("true")) // true

	// mengubah string ke float
	fmt.Println(strconv.ParseFloat("3.14", 64)) // 3.14 // 64 adalah bit size

	// mengubah string ke int64
	fmt.Println(strconv.ParseInt("100", 10, 64)) // 100 // 10 adalah base // 64 adalah bit size

	// mengubah bool ke string
	fmt.Println(strconv.FormatBool(true)) // true

	// mengubah float ke string
	fmt.Println(strconv.FormatFloat(3.14, 'f', -1, 64)) // 3.14 // 'f' adalah format // -1 adalah precision // 64 adalah bit size
}