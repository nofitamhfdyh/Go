package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string, umur int) {
	firstName = "Nofita"
	middleName = "Mahfudiyah"
	lastName = "Nupnup"
	umur = 23

	return
}
func main() {
	firstName1, middleName1, lastName1, umur1 := getCompleteName()
	fmt.Println(firstName1, middleName1, lastName1)
	fmt.Println("Dia berumur", umur1)
}
