package main

import "fmt"

func random() interface{} {
	return "ok"
}

// func main() {
// 	// type assertions => merupakan cara untuk mengubah tipe data dari sebuah interface
// 	var result interface{} = random()
// 	var resultString string = result.(string)
// 	fmt.Println(resultString)
// }