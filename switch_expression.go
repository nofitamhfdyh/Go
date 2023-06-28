package main

import "fmt"

func main() {
	name := "Jay"

	switch name {
	case "Jay":
		fmt.Println("Hello Jay")
	case "Mahfudiyay":
		fmt.Println("Hello Mahfudiyay")
	default:
		fmt.Println("HI, Boleh kenalan?")
	}

	//short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Benar")
	}

	//switch tanpa kondisi dan variabel
	length1 := len(name)
	switch { //tidak perlu menambah kondisi setelah switch, kondisi bisa ditambahkan setelah 'case' -> hampir mirip dengan if else expression
	case length1 > 10:
		fmt.Println("Nama Terlalu Panjang")
	case length1 > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Sudah benar")

	}
}
