package main

import "fmt"

func getFullName() (string, string, int) { //harus disebutkan semua tipe data returnnya, walaupun sama, ttp harus disebutkan dan urut
	return "Nofita", "Mahfudiyah", 23
}
func main() {
	firstName, lastName, umur := getFullName()
	fmt.Println(firstName, lastName, "berumur", umur)

	//jika tidak ingin memanggil seluruh return value / ada beberapa return value yang tidak dipakai, maka bisa menggunakan tanda underscore / _
	firstName1, _, umur1 := getFullName()
	fmt.Println(firstName1, "berumur", umur1)
}
