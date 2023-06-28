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

	//append slice
	days := []string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}
	daysSlice1 := days[5:]
	daysSlice1[0] = "Sabtu Baru"
	daysSlice1[1] = "Minggu Baru"
	fmt.Println(days) // output -> [Senin Selasa Rabu Kamis Jumat Sabtu Baru Minggu Baru]

	fmt.Println(daysSlice1) // output -> [Sabtu Baru Minggu Baru]

	daySlice2 := append(daysSlice1, "Libur Baru") // append -> utk membuat slice baru dengan menambahkan data di posisi terkakhir slice sebelumnya. jika kapasitas array sebelumnya sudah penuh, maka akan membuat array baru
	daySlice2[0] = "Ups"
	fmt.Println(daySlice2) // output -> [Ups Minggu Baru Libur Baru]
	fmt.Println(days)      // output -> [Senin Selasa Rabu Kamis Jumat Sabtu Baru Minggu Baru]

	//membuat slice dengan func make
	newSLice := make([]string, 2, 5) // panjang nya 2, kapasitasnya 5
	newSLice[0] = "Nofita"
	newSLice[1] = "Mahfudiyah"

	fmt.Println(newSLice)
	fmt.Println(len(newSLice))
	fmt.Println(cap(newSLice))

	//copy slice
	fromSlice := days[:]
	toSLice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSLice, fromSlice) // destinasinya adalah toSLice, datanya dari fromSLice
	fmt.Println(toSLice)

	//perbedaa array dengan slice
	iniArray := [...]int{1, 2, 3, 4, 5} // kalo mau buat array, pastikan ada keterangan di kurung depan. jika tidak ada batasan panjang array, maka pakai titik 3 dalam kurung
	iniArray1 := [2]int{1, 2}           //keterangan panjang array
	iniSLice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniArray1)
	fmt.Println(iniSLice)
}
