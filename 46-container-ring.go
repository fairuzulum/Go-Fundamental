package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	// container ring adalah package yang digunakan untuk membuat ring
	// ring adalah list yang memiliki pointer ke elemen pertama
	data := ring.New(5)

	for i := 0; i < 5; i++ {
		data.Value = "value: " + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	data.Do(func(value interface{}) {
		fmt.Println(value)
	}) // do adalah method yang digunakan untuk melakukan iterasi pada ring
}