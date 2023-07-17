package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Nofita Mahfudiyah", "Nofita")) //apakah ada kata 'nofita' dalam 'nofita mahfudiyah' -> output true / false
	fmt.Println(strings.Contains("Nofita Mahfudiyah", "Nupnup"))

	fmt.Println(strings.Split("Nofita Mahfudiyah", " ")) //memotong string menjadi slice

	fmt.Println(strings.ToLower("Nofita Mahfudiyah"))
	fmt.Println(strings.ToUpper("Nofita Mahfudiyah"))
	fmt.Println(strings.ToTitle("nofita mahfudiyah"))

	fmt.Println(strings.Trim("         Nofita Mahfudiyah       ", " ")) //menghilangkan karakter sesuai cutset di paling kiri dan paling kanan
	fmt.Println(strings.Trim("a     Nofita Mahfudiyah      a", " "))

	fmt.Println(strings.ReplaceAll("Nofita Mahfudiyah Nofita Nupnup", "Nofita", "Nupnup"))
}
