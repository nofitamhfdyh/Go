package main

import "fmt"

func main() {
	//break digunakn untuk menghentikan seluruh perulangan
	//continue digunakan untuk menghentikan perulangan yang saat itu sedang berjalan, tapi langsung melanjutkan ke perulangan selanjutnya

	//break
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Perulangan ke : ", i)
	}

	//continue
	for i := 0; i < 10; i++ {
		if i%2 == 0 { // jika i modul / dibagi 2 = 0
			continue
		}
		fmt.Println("Perulangan ke -> ", i)
	}
}
