package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Put Your Database Host") //tipe datanya bisa diubah, tidak harus string
	username := flag.String("username", "root", "Put Your Database Username")
	password := flag.String("password", "root", "Put Your Database Password")

	flag.Parse() //kalau parse ini tidak dipanggil, maka outputnya akan sesuai dengan default value nya, bukan sesuai inputan dari user (1)

	fmt.Println(*host, *username, *password) // terminal -> go run flag.go ; atau ditambah dengan -host=nupnup -username=mahfud -password=yuhu (1)

	fmt.Println("Host :", *host) //kalau tidak diberi tanda *, maka outputnya tidak sesuai dengan value yang diinginkan
	fmt.Println("Username :", *username)
	fmt.Println("Password :", *password)
}
