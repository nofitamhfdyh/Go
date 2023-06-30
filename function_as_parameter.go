package main

import "fmt"

func sayHelloWithFilter(name string, filter func(string) string) {
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
