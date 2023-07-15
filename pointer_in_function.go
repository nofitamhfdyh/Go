package main

import "fmt"

type Alamat struct {
	Kota, Provinsi, Negara string
}

func ChangeAddressToIndonesia(address *Alamat) {
	address.Negara = "Indonesia"
}
func main() {
	// secara default -> pass by value
	// maka, jika mengubah data dlm function, data asli tidak akan pernah berubah -> variable nya aman krn tdk bisa dirubah
	// namun saat tertentu, kita ingin membuat function yang bisa mengubah data asli parameter tsb -> dilakukan pointer di function
	// utk menjadikan sebuah parameter sbg pointer, bisa menggunakan operator * di parameternya

	addres := Alamat{
		Kota:     "Blitar",
		Provinsi: "Jawa Timur",
		Negara:   "",
	}
	//cara memanggil 1
	ChangeAddressToIndonesia(&addres)

	fmt.Println(addres) // tidak berubah, karena yang dikirim adalah data copyan value nya

	//cara memanggil 2
	var alamatPointer *Alamat = &addres
	ChangeAddressToIndonesia(alamatPointer)
	fmt.Println(alamatPointer)
}
