package main

import "fmt"

func main() {
	person := map[string]string{ // [string] -> tipe key ; string{} -> tipe value
		"name":    "Nofita",
		"address": "Blitar",
	}

	person["title"] = "Technical Writer" // nambah data di map

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)
	book["title"] = "Buku Go-Lang"
	book["author"] = "NOfita Mahfudiyah"
	book["wrong"] = "Ups"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "wrong")
	fmt.Println(book)
	fmt.Println(len(book))

	//jika menggunakan perulangan for
	biodata := make(map[string]string)
	biodata["Nama"] = "Nofita"
	biodata["Pekerjaan"] = "Programmer"
	biodata["Alamat"] = "Blitar"

	for key, value := range biodata {
		fmt.Println(key, " = ", value)
	}

}
