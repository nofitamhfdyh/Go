package helper

import "fmt"

var version = 1 // (3) tidak bisa diakses krn diawali dgn huruf kecil
var Application = "Belajar Golang"

func SayHello(name string) {
	fmt.Println("Hello", name)
}

func sayGoodBye(name string) { // jika function diawali dgn huruf kecil, maka tidak bisa diakses di luar package (2)
	fmt.Println("Hallo", name)
}
