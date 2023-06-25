package main

import "fmt"

func main() {

	fmt.Println("Hello World!")
	fmt.Println("Anjaaz kelazz!")

	fmt.Println("Satu = ", 1)
	fmt.Println("Dua", 2)
	fmt.Println("Tiga Koma Lima", 3.5)

	fmt.Println("Benar :", true)
	fmt.Println("Salah :", false)

	fmt.Println(len("Nofita"))          //untuk menghitung jumlah karakter dari string
	fmt.Println("Nofita Mahfudiyah"[0]) //hasilnya adalah representasi dari byte huruf N
	fmt.Println("Nofita Mahfudiyah"[1]) //hasilnya adalah representasi dari byte huruf O

	//(1)saat membuat variabel, wajib menuliskan tipe data dari variable
	var name string

	name = "Nofita Mahfudiyah"
	fmt.Println(name)

	name = "Farhan Rizky Alkarim" //value variable bisa diubah dlm golang
	fmt.Println(name)

	//(2)jika langsung menginisialisasi data pada variable, tidak wajib menuliskan tipe data
	var nama = "Nupnup" //sudah diidentifikasi sendiri dengan golang tipe data berdasarkan value yang diinputkan
	fmt.Println(nama)

	var age = 45
	fmt.Println(age)

	country := "Indonesia" //kata kunci 'var' tidak wajib, jika langsung diinisialisasi datanya -> namun memakai := pada awal inisialisasi
	fmt.Println(country)

	country = "Nopnop" //jika mau dipakai lagi, cukup pakai =, tidak perlu :=
	fmt.Println(country)

	//deklarasi nultiple variable
	var (
		firstName = "Nofita"
		lastName  = "Mahfudiyah"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	//constant
	const namaDepan string = "Nofita"
	const namaBelakang = "Mahfudiyah"
	//fmt.Println(namaDepan) --> walaupun const tidak dipakai, program akan tetap berjalan. beda dengan var
	fmt.Println(namaBelakang)

	//error constant -> value dari variabel pertama tidak bisa diubah jika memakai constant, beda dengan var
	//namaDepan = "Nupnup"
	//namaBelakang = "Nopnop"

	//deklarasi multiple constant
	const (
		firstNama string = "Nofita"
		lastNama         = "Mahfudiyah"
		value            = 1000
	)

	fmt.Println(firstNama)
	fmt.Println(lastNama)
	fmt.Println(value)
}
