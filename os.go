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

}
