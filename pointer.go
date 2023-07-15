package main

import "fmt"

type Address struct {
	Kecamatan, Kota, Provinsi string
}

func main() {
	//(1)Pass by Value -> jika kita mengirim variabel ke sebuah function / method / variable lain, yang dikirim adalah duplikasi valuenya. sudah by default dari Golang

	//(1)
	address1 := Address{
		Kecamatan: "Kademangan",
		Kota:      "Blitar",
		Provinsi:  "Jawa Timur",
	}
	address2 := address1

	address2.Kota = "Surabaya"

	fmt.Println(address1) //tidak berubah isinya, tetap -> ini yang dimaksud dengan pass by value (1)
	fmt.Println(address2)
}
