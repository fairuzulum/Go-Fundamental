package main

import "fmt"

func tipeDataMap() {
	// tipe data Map adalah tipe data yang memiliki kumpulan data dengan key dan value

	// deklarasi map
	person := map[string]string{
		"name":    "fairuz",
		"country": "Indonesia",
	}
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["country"])

	// menambahkan data ke dalam map
	person["city"] = "Jakarta"
	// merubah data map
	person["name"] = "ulumu"
	fmt.Println(person["city"])
	fmt.Println(person["name"])


	// len() => untuk mengetahui panjang map
	fmt.Println(len(person))

	// delete(map, key) => untuk menghapus data dari map
	delete(person, "city")
	fmt.Println(person)

	// make(map[tipe data key]tipe data value) => membuat map baru
	newPerson := make(map[string]string)
	newPerson["name"] = "fairuz"
	newPerson["country"] = "Indonesia"
	fmt.Println(newPerson)

}