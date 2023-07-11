package main

type Customer struct { //utk struct, rata" namanya diawalai dengan huruf kapital. kalau ada 2 kata, brti 2 huruf paling depan kapital
	Nama, Alamat string
	Umur         int
}

func main() {
	// mmenggabungkan nol / lebih data yang tipenya beda dalam 1 kesatuan.
	var nofita Customer
	nofita.Nama = "Nofita Mahfudiyah"
	nofita.Alamat = "Blitar"
	nofita.Umur
}
