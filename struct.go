package main

import "fmt"

// (1)
type Customer struct { //utk struct, rata" namanya diawalai dengan huruf kapital. kalau ada 2 kata, brti 2 huruf paling depan kapital
	Nama, Alamat string
	Umur         int
	//Status       bool
}

// struct method(1)
func (customer Customer) sayStruct(name string) {
	fmt.Println("Hello ", name, "My Name is ", customer.Nama)
}

func (aaa Customer) sayAAA() {
	fmt.Println("Say huuu", aaa.Nama)
}

func main() {
	// mmenggabungkan nol / lebih data yang tipenya beda dalam 1 kesatuan.
	// sangat cocok utk merepresentasikan data daripada memakai map atau array, karena strukturnya lebih bagus
	//(1)
	var nupnup Customer
	nupnup.Nama = "Nofita Mahfudiyah"
	nupnup.Alamat = "Blitar"
	nupnup.Umur = 23

	fmt.Println(nupnup)

	fmt.Println(nupnup.Nama)
	fmt.Println(nupnup.Alamat)
	fmt.Println(nupnup.Umur)

	//(2) struct literals
	Farhan := Customer{
		Nama:   "Farhan Rizky Alkarim",
		Alamat: "Tulungagung",
		Umur:   17,
	}

	fmt.Println(Farhan)

	//(3) struct literals
	Pandan := Customer{"Pandan Zahwa Azizah", "Sukoharjo, Solo", 23} // bisa seperti ini yang penting urutan sesuai. kelemahannya, ketika didalam struct ditambah value baru, dia akan otomatis error
	//error juga ketika urutan deklarasi struct dibalik
	fmt.Println(Pandan)

	//struct method (1)
	// sayStruct(nupnup, "Nopnop") //pemanggilan tanpa tanpa struct
	nupnup.sayStruct("Nopnop ")

	nupnup.sayAAA()
}
