package main

import "fmt"

func main() {
	// closure = kemampuan function berinteraksi dengan data yang ada didalamnya
	nama := "Nofita"
	counter := 0
	increment := func() {
		nama = "Mahfudiyah" // variabel yang diinisialisasi diatas, bisa dipanggil dibawahnya, tapi kalo dibalik tidak bisa
		fmt.Println("Increment")
		counter++

	}

	increment()
	increment()
	fmt.Println(nama)    // is dari variabel ini diganti di line 10
	fmt.Println(counter) // isi dari variabel ini diganti di line 12
}
