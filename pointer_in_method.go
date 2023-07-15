package main

import "fmt"

type Man struct {
	Nama string
}

func (man *Man) Married() {
	man.Nama = "Ms. " + man.Nama
	//fmt.Println("DI Method", man.Nama)
}
func main() {
	// walaupun method akan menempel di struct, sebenarnya data struct yang diakses dalam method adalah pass by value
	// sehingga, direkomendasikan menggunakan pointer di method -> agar tudak boros memory

	nupnup := Man{"Nofita"}
	nupnup.Married()

	fmt.Println(nupnup.Nama)
}
