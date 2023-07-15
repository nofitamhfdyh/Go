package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) { //kalau utk cek error, biasanya return value nya ada 2, yg 1 nilai aslinya dan 1 nya adalah nilai errornya
	if pembagi == 0 {
		return 0, errors.New("Pembagi Dengan Nol") //mmebuat error
	} else {
		result := nilai / pembagi
		return result, nil //penggunaan nil sebagai return dari error,
	}
}

func main() {
	//go lang memiliki interface yang digunakan sbg kontrak untuk membuat error, nama interfacenya adalah error
	//var cthError error = errors.New("Ups Error")
	//fmt.Println(cthError.Error())
	hasil, err := Pembagian(100, 10)
	if err == nil {
		fmt.Println("Hasil ", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}
