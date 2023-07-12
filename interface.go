package main

import "fmt"

type HashName interface { //nama kontraknya HashName
	GetName() string //method getname dengan returnnya string
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func SayHello(hashName HashName) { //interface juga bisa digunakan sebagai parameter di func
	fmt.Println("Hello", hashName.GetName())
}
func main() {
	//didalam interface, isinya hanya definisi method saja, only untuk kontrak
	person := Person{Name: "Nofita Mahfudiyah"}
	SayHello(person)
}
