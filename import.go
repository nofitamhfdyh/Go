package main

import (
	"fmt"
	"test123/helper"
)

func main() {
	helper.SayHello("Nofita")
	// run output di terminal -> go run import.go

	// helper.sayGoodBye("Nofita") -> jika function diawali dgn huruf kecil, maka tidak bisa diakses di luar package (2)
	fmt.Println(helper.Application)
	//fmt.Println(helper.version) -> (3) tidak bisa diakses krn diawali dgn huruf kecil

}
