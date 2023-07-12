package main

import "fmt"

type Customer struct { //utk struct, rata" namanya diawalai dengan huruf kapital. kalau ada 2 kata, brti 2 huruf paling depan kapital
	Nama, Alamat string
	Umur         int
}

func main() {
	// mmenggabungkan nol / lebih data yang tipenya beda dalam 1 kesatuan.
	// sangat cocok utk merepresentasikan data daripada memakai map atau array, karena strukturnya lebih bagus
	var me Customer
	me.Nama = "Nofita Mahfudiyah"
	me.Alamat = "Blitar"
	me.Umur = 23

	fmt.Println(me)

	fmt.Println(me.Nama)
	fmt.Println(me.Alamat)
	fmt.Println(me.Umur)
}
