package main

import "fmt"

func random() interface{} { //returnya adalah interface kosong
	return "OK"
}

func main() {
	//merubah tipe data sesuai dengan yang kita inginkan, digunakan saat bertemu dengan interface kosong

	result := random()
	resultString := result.(string) // melakukan konversi dari interface kosong menjadi string
	fmt.Println(resultString)

	//panic, cth konversi yang tidak sesuai
	//resultInt := result.(int)
	//fmt.Println(resultInt)
}
