package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}
func main() {
	// data nill adalah data kosong
	// hanya bisa utk beberapa tipe data, yaitu interface, function, map, slice, pointer dan channel
	fmt.Println(newMap("Nofita"))

	person := newMap("Mahfudiyah")
	fmt.Println(person)

	//penggunaan nill
	data := newMap("")
	if data == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(data)
	}
}
