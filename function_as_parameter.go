package main

import "fmt"

type Filter func(string) string // utk membuat alias dari funtion, supaya tidak terlalu panjang ketika functionnya dipakai sbg parameter di function lain

func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println("Hello ", filter(name))
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	//function juga bisa digunakan sbg parameter utk funtion lain
	//(1)
	sayHelloWithFilter("Anjing", spamFilter)
	//(2)
	filter := spamFilter
	sayHelloWithFilter("Nupnup", filter)
}
