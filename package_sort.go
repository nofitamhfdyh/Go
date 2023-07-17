package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User //alias utk UserSlice -> array dari []User ; dan mau mengurutkan data di UserSLice

//membuat kontrak untuk interface sort dari len(), less(), swap()

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age //sorting nya berdasarkan umur (1)
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}
func main() {
	//agar data bisa diurutkan, maka harus mengimplementaiskan kontrak di interdace sort.Interface

	//proses sorting
	users := []User{
		{"Nofita", 20},
		{"Mahfudiyah", 40},
		{"Farhan", 23},
		{"Rizky", 59},
		{"Alkarim", 9},
	}

	fmt.Println(users) //sebelum di sort umur paling muda (1)

	sort.Sort(UserSlice(users))
	fmt.Println(users) //setelah di sort
}
