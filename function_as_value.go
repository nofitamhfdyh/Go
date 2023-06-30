package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}
func main() {
	//dlm golang, function bisa disimpan dalam variable dan  bisa dianggap sbg value
	goodBye := getGoodBye //getGoodBye di set sebagai value dan bisa dimasukkan dalam variable goodBye
	//(1)
	fmt.Println(goodBye("Nofita"))
	//(2)
	result := goodBye("Nupnup")
	fmt.Println(result)
}
