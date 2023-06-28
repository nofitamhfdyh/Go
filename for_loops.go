package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke : ", counter)
		counter++
	}

	// for dengan statement -> untuk mempersingkat kode yang diatas
	//(1) init statement -> dieksekusi sebelum for dieksekusi
	//(2)post statement -> dieksekusi diakhir tiap perulangan

	for counter1 := 1; counter1 <= 10; counter1++ { // counter1 :=1 -> init statement ; counter1++ -> post statement
		fmt.Println("Perulangan ke -> ", counter1)
	}

	//for range -> bisa dugunakan di data collection
	slice := []string{"Nofita", "Mahfudiyah", "Jay"}

	//menggunakan cara manual
	//fmt.Println(slice[0])
	//fmt.Println(slice[1])
	//fmt.Println(slice[2])

	//menggunakan perulangan tapi masih secara manual
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	//menggunakan for range
	for i, value := range slice {
		fmt.Println("Index ke ", i, ": ", value)
	}

	//jika ingin menggunakan for range tapi tanpa menampilkan index, jadi hanya value saja, maka bisa diganti dengan underscore / _ supaya tidak error
	for _, value := range slice {
		fmt.Println(value)
	}

	//bisa digunakan di map
	person := make(map[string]string)
	person["nama"] = "Nofita"
	person["posisi"] = "technical writer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
