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

	//saat mengubah variabel pointer, berarti yang berubah hanya variabel
	// seluruh variabel yang mengacu ke data yang sama tidak akan berubah
	// jika ingin mengubah seluruh variable yang mengacu ke data tsb, menggunakan --> operator * (3)

	// sebelum membuat pointer menggunakan operator &, golang memiliki function new -> digunakan untuk membuat pointer
	// tapi, function new (4) hanya mengembalikan pointer ke data kosong -> jadi tidak ada data awal

	//(1)
	address1 := Address{
		Kecamatan: "Kademangan",
		Kota:      "Blitar",
		Provinsi:  "Jawa Timur",
	}
	//address2 := address1 // (1)
	address2 := &address1
	address3 := &address1

	address2.Kota = "Surabaya"

	//(3)
	*address2 = Address{ //tambah bintang didepan variabel, sehingga address 1 dipaksa untuk mengikuti isi dari address 2
		Kecamatan: "Pasar Minggu",
		Kota:      "Jakarta Selatan",
		Provinsi:  "DKI Jakarta",
	}

	fmt.Println(address1) //tidak berubah isinya, tetap -> ini yang dimaksud dengan pass by value (1)
	fmt.Println(address2) //didepannya jadi ada & -> memberitahu bahwa kalau address 2 itu pointer -> cara cek nya klik ctrl + spasi
	fmt.Println(address3)

	//(4)
	address4 := new(Address)

	address4.Kota = "Jogja"

	fmt.Println(address4)
}
