package main

import (
	"fmt"
)

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	defer logging() // setelah function runApplication dieksekusi, func Logging pasti dieksekusi dipaling akhir walaupun penempatannya diawal, karena pakai defer. tpi pemakaian defer diusahakan diawal, jadi ketika didalam func terjadi error, tidak kena dampak
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println("Result ", result)
}

func endApp() {
	message := recover()
	//(3)
	if message != nil {
		fmt.Println("Error dengan message", message)
	}
	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("ERROR")
	}

	//(1) dipindahkan ke func defer function
	//message := recover()
	//fmt.Println("Error dengan message", message) // outputnya tidak ada, karena berada di dalam 1 func yang error dan dibawah kode progam yang error. jadi tidak akan dieksekusi
	fmt.Println("Aplikasi Berjalan")
}
func main() {
	//defer func = penjadwalan eksekusi function setelah function lain dieksekusi. akan tetap dieksekusi walaupun func yg dieksekusi sebelumnya error
	//panic func = digunakan utk menghentikan program dan dipakai ketika terjadi error saat program sedang berjalan. namun,ketika func panic dipanggil,program berhenti, tapi ketika menggunakan defer, maka defer func akan tetap dieksekusi
	// recover func = digunakan utk menangkap panic. proses panic akan berhenti, shg program ttp berjalan

	runApplication(10)
	runApp(true)       //outputnya akan nil jika isinya false dan tidak menggunakan if else (3)
	fmt.Println("Eko") // utk cek, ketika func error, apakah aplikasi akan berhenti. output eko ttp keluar, jadi dipastikan kalau aplikasi teteap berjalan
}
