package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Round(1.7)) //pembulatan paling dekat keatas
	fmt.Println(math.Round(1.3)) //pembulatan paling dekat kebawah
	fmt.Println(math.Floor(1.7)) //dipaksa bulat ke bawah
	fmt.Println(math.Ceil(1.3))  //dipaksa bulat keatas

	fmt.Println(math.Max(10, 20))
	fmt.Println(math.Min(10, 20))
}
