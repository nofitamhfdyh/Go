package main

import (
	"fmt"
)

func main() {
	names := []string{"Nofita", "Mahfudiyah", "Farhan", "Rizky", "Alkarim", "Nupnup"}
	slice := names[4:6]

	fmt.Println(slice[0])
	fmt.Println(slice[1])

	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1)) //len utk mendapatkan panjang
	fmt.Println(cap(slice1)) //cap untuk mendapatkan kapasitas

	//months[5] = "Diubah"
	//fmt.Println(slice1)

	slice1[0] = "Mei Lagi" // index 0 tidak mengacu ke index array months, tapi indek 0 dari slice yang dimulai dari index ke-4 dari array
	fmt.Println(months)

}
