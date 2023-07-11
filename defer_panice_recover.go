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
func main() {
	//defer func = penjadwalan eksekusi function setelah function lain dieksekusi. akan tetap dieksekusi walaupun func yg dieksekusi sebelumnya error
	runApplication(10)
}
