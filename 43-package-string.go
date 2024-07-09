package main

import (
	"fmt"
	"strings"
)

func main() {
	// package string adalah package yang digunakan untuk memanipulasi string

	fmt.Println(strings.Contains("Hello World", "World")) // true
	fmt.Println(strings.Split("Fairuz Ulum", " ")) // ["Fairuz", "Ulum"]
	fmt.Println(strings.ToLower("HELLO WORLD")) // hello world
	fmt.Println(strings.ToUpper("hello world")) // HELLO WORLD
	fmt.Println(strings.Trim("   Hello World   ", " ")) // Hello World
	fmt.Println(strings.ReplaceAll("Hello World", "World", "Fairuz Ulum")) // Hello Fairuz Ulum
	fmt.Println(strings.Compare("Hello World", "Hello World")) // 0
	fmt.Println(strings.ContainsAny("Hello World", "abc")) // false
	fmt.Println(strings.Count("Hello World", "l")) // 3
	fmt.Println(strings.Fields("Hello World")) // ["Hello", "World"]
	fmt.Println(strings.HasPrefix("Hello World", "Hello")) // true
	fmt.Println(strings.HasSuffix("Hello World", "World")) // true
	fmt.Println(strings.Index("Hello World", "o")) // 4
	fmt.Println(strings.Join([]string{"Hello", "World"}, " ")) // Hello World
	fmt.Println(strings.Repeat("Hello World ", 3)) // Hello World Hello World Hello World
	
}