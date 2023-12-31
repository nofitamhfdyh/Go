package main

import "fmt"

func main() {
	var result = 10 + 10
	fmt.Println(result)

	var a = 10
	var b = 10
	var c = a * b
	fmt.Println(c)

	var i = 10
	i += 10 // i = i + 10
	fmt.Println(i)

	i++ // i = i + 1
	fmt.Println(i)

	var negative = -100
	var positive = +100
	fmt.Println(negative)
	fmt.Println(positive)

	//operasi perbandingan
	var name1 = "Zoya"
	var name2 = "Zoya"

	var result1 bool = name1 == name2

	fmt.Println(result1)

	var value1 = 100
	var value2 = 200
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)

	//Boolean Operation
	var nilaiAkhir = 90
	var nilaiAbsensi = 80

	//(1)cara 1
	//var lulusNilaiAkhir bool = nilaiAkhir > 80
	//var lulusNilaiAbsensi bool = nilaiAbsensi > 80
	//fmt.Println(lulusNilaiAkhir)
	//fmt.Println(lulusNilaiAbsensi)
	//
	//var lulus bool = lulusNilaiAkhir && lulusNilaiAbsensi
	//fmt.Println(lulus)

	//(2) cara 2
	fmt.Println(nilaiAkhir >= 80 && nilaiAbsensi >= 80)
}
