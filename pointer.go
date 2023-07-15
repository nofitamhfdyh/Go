package main

import "fmt"

type Address struct {
	Kecamatan, Kota, Provinsi string
}

func main() {
	//(1)Pass by Value -> jika kita mengirim variabel ke sebuah function / method / variable lain, yang dikirim adalah duplikasi valuenya. sudah by default dari Golang
	//konsep di bahasa pemrograman lain menggunakan pass by referemce
	// Pointer -> kemampuan yang digunakan utk membuat reference ke lokasi data di memory yang sama tanpa mendupilkasi data yang sudah ada.
	// membuat pass by reference menggunakan pointer
	// cara membuat pointer itu menggunakan operator & (and)
	// (2) operator & -> membuat variable dengan nilai pointer ke variable yang lain -> operator & diikuti dengan nama variable nya

	//(1)
	address1 := Address{
		Kecamatan: "Kademangan",
		Kota:      "Blitar",
		Provinsi:  "Jawa Timur",
	}
	//address2 := address1 // (1)
	address2 := &address1

	address2.Kota = "Surabaya"

	fmt.Println(address1) //tidak berubah isinya, tetap -> ini yang dimaksud dengan pass by value (1)
	fmt.Println(address2) //didepannya jadi ada & -> memberitahu bahwa kalau address 2 itu pointer -> cara cek nya klik ctrl + spasi
}
