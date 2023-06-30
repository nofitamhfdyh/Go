package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blacklist BlackList) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

//(1)
//func blackListAdmin(name string) bool  {
//	return name == "Admin"
//}
//
//func blackListRoot(name string) bool  {
//	return name == "Root"
//}

func main() {
	//(1) dideklarasikan dulu dalam variable
	blackList := func(name string) bool {
		return name == "Anjing"
	}
	registerUser("Nofita", blackList)
	registerUser("Anjing", blackList)

	//(2) langsung dimasukkan sbg parameter
	registerUser("Anjing", func(name string) bool {
		return name == "Anjing"
	})
}
