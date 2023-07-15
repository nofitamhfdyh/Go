package main

import "fmt"

func random() interface{} { //returnya adalah interface kosong
	return 100
}

func main() {
	//merubah tipe data sesuai dengan yang kita inginkan, digunakan saat bertemu dengan interface kosong
	//agar mengurangi risiko panic dalam type assertions, maka digunakan penggunaan switch dalam implementasinya

	result := random()
	//resultString := result.(string) // melakukan konversi dari interface kosong menjadi string
	//fmt.Println(resultString)
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown Type")
	}

	//panic, cth konversi yang tidak sesuai
	//resultInt := result.(int)
	//fmt.Println(resultInt)
}
