package main

import (
	"fmt"
	"strconv"
)

func main() {
	// format -> dari .....(tipe data lain) menjadi String
	// Parse -> dari Strimg menjadi ... (tipe data tujuan)

	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println("Error", err.Error())
	}

	number, err := strconv.ParseInt("1000000", 10, 64) //base : jenis base angka -> binary=2 ; decimal=10 ; oktal=8 ; hexadecimal=16
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println("Error", err.Error())
	}

	value := strconv.FormatInt(1000000, 10)
	fmt.Println(value)

	//(1) kalau ingin melihat errornya
	valueInt, err := strconv.Atoi("100000") // atoi = conversi simple dari string ke int (tanpa harus menambah base dan bit size, jadi otomatis menggunakan base 10 / decimal); itoa = conversi simple dari int ke string
	if err == nil {
		fmt.Println(valueInt)
	} else {
		fmt.Println("Error", err.Error())
	}

	//(2) kalau ingin mengabaikan errornya
	valInt, _ := strconv.Atoi("35000")
	fmt.Println(valInt)
}
