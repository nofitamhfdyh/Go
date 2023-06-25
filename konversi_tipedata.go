package main

import (
	"fmt"
)

func main() {
	var nilai32 int32 = 100000 // kalau nilainya 127, output nya sama semua. tapi jika 128, outputnya nilai8 akan menjadi -128 (balik lagi ke titik bawah dari int8)
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32) //terjadi integer overflow -> saat sudah mencapai batas, balik lagi ke titik yang paling bawah

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var name = "Nofita"
	var e byte = name[0] // disebut byte atau alias dari uint8

	var eString string = string(e) //konversi dari byte menjadi string

	fmt.Println(name)
	fmt.Println(eString)
}
