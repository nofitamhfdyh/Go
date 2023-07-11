package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	defer logging() // setelah function runApplication dieksekusi, func Logging pasti dieksekusi dipaling akhir walaupun penempatannya diawal, karena pakai defer. tpi pemakaian defer diusahakan diawal, jadi ketika didalam func terjadi error, tidak kena dampak
	fmt.Println("Run Application")
}
func main() {
	//defer func = penjadwalan eksekusi function setelah function lain dieksekusi. akan tetap dieksekusi walaupun func yg dieksekusi sebelumnya error
	runApplication()
}
