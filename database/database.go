package database

var connection string

/**
(1)
- Saat membuat package, kita bisa membuat sebuah function yang akan diakses ketika package kita diakses
- hal ini sangat cocok ketika package kita berisi function" yang berkomunikasi dengan database, kita membuat function inisialisasi untuk membuka koneksi ke database
- untuk membuat function yang diakses secara otomatis ketika package diakses, kita cukup membuat function dengan nama init (1)
*/

func init() { //(1)
	// fmt.Println("Init Sukses Dipanggil") //untuk memastikan penggunaan blank identifier berhasil
	connection = "MYSQL"
}

func GetDatabase() string {
	return connection
}
