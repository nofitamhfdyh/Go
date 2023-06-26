package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Nofita"
	names[1] = "Mahfudiyah"
	names[2] = "Nupnup"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	//membuat array secara langsung
	var values = [3]int{
		90, 80, 70,
		//90,
		//80,
		//70,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	values[0] = 100 //mengubah data diposisi index
	fmt.Println(values)

	//function dalam array
	fmt.Println(len(names)) //utk mengetahui jumlah panjang arraynya, bukan jumlah datanya
	fmt.Println(len(values))

	var ceklagi [10]string
	fmt.Println(len(ceklagi)) //output 10, walaupun blm ada datanya, tapi jumlah panjang arraynya sudah terhitung 10
}
