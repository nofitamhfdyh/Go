package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	firstName := "Nofita"
	sayHelloTo(firstName, "Mahfudiyah") //urutan parameter harus sama sesuai deklarasi di func induk
	sayHelloTo("Jayana", "Citra Agung Pramu Putra")
	sayHelloTo("Nupnup", "Nopnop")
}
