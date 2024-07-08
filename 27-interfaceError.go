package main

import (
	"errors"
	"fmt"
)

func pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		return nilai / pembagi, nil
	}
}

func errorInterface() {
	// interface error
	var contohError error = errors.New("Opps Error")
	fmt.Println(contohError.Error())

	result, error := pembagi(100, 5)
	if error != nil {
		fmt.Println("Error:", error)
	} else {
		fmt.Println("Hasil:", result)
	}

}