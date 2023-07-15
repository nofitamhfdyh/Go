package main

/**
Blank Identifier
- Kadang kita hanya ingin menjalankan init function di package tanpa harus mengeksekusi salah satu function yang ada di package
- Secara default, Golang akan komplain ketika ada package yang diimport namun tidak digunakan
- Sehingga kita bisa menggunakan -> blank identifier -> (_) -> penempatannya didepan nama package ketika melakukan import
*/

import _ "test123/database" //blank identifier (2)

func main() {
	//result := database.GetDatabase()
	//fmt.Println(result)
}
