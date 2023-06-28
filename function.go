package main

import "fmt"

func sayHello() {
	fmt.Println("Hello")
}
func main() {
	//func -> kode yang dibuat agar bisa digunakan secara berulang
	//func bisa dipanggil dengan nama func nya diikuti dengan tanda kurung buka dan kurung tutup
	sayHello()
	for i := 0; i < 4; i++ {
		sayHello()
	}
}
