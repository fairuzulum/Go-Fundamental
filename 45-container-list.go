package main

import (
	"container/list"
	"fmt"
)

func main() {
	// container list adalah package yang digunakan untuk membuat list

	data := list.New()

	data.PushBack("Ahmad")
	data.PushBack("Fairuz")
	data.PushBack("Ulum")

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}