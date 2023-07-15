package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Argument :")
	fmt.Println(args) //lokasi file arguments

	// cara mengakses -> terminal -> go run os.go nofita mahfudiyah
	fmt.Println(args[1]) // output nofita
	fmt.Println(args[2]) // output mahfudiyah

	//bisa buka di link documentation -> golang.org/pkg/os/
	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname : ", hostname) // terminal -> go run os.go nupnup nupnup
	} else {
		fmt.Println("Error", err.Error())
	}

	//get environtment utk koneksi ke database lokal
	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")

	fmt.Println(username)
	fmt.Println(password)
	// terminal -> export APP_USERNAME=root
	// terminal -> export APP_PASSWORD=root
	// utk cek bisa langsung ke -> go run os.go
}
