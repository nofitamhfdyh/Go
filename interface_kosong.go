package main

import "fmt"

func Ups(i int) interface{} { //dalam interface tidak menyebutkan tipe balikannya, brti seluruh tipe data bisa
	//return 1
	//return true
	//return "Ups"

	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func main() {
	//interface yang tidak memiliki method -> sehingga secara otomatis seluruh tipe data akan menjadi implementasinya
	/**
	contoh penggunaan :
	- fmt.Println (a... interface{})
	- panic(v interface{})
	- recover() interface{}
	dll
	-
	*/

	kosong := Ups(2)
	fmt.Println(kosong)
}
