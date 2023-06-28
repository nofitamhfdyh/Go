package main

import "fmt"

func main() {
	name := "mahfudiyah"

	if name == "Eko" {
		fmt.Println("Hello Eko")
	} else if name == "Joko" {
		fmt.Println("Hello Joko")
	} else {
		fmt.Println("Hi, Boleh kenalan?")
	}

	//short statement
	//var length = len(name) -> bisa digabung di bawah setelah 'if'
	if length := len(name); length > 5 { //len() juga bisa digunakan utk cek panjang string
		fmt.Println("Nama Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}
}
