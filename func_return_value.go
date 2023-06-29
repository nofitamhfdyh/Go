package main

import "fmt"

func getHello(name string) string {
	return "Hello " + name
}
func getHai(nama string) string {
	if nama == "" {
		return "Hello Bro"
	} else {
		return "Hello " + nama
	}
}
func main() {
	result := getHello("Nofita")
	fmt.Println(result)

	//(1)
	resultHai := getHai("")
	fmt.Println(resultHai)

	//(2) cara singkat
	fmt.Println(getHai("Nupnup"))
}
