package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Nofita")
	data.PushBack("Mahfudiyah")
	data.PushBack("Nupnup")
	data.PushBack("Farhan")
	data.PushBack("Rizky")
	data.PushBack("Alkarim")
	data.PushFront("Noooooppppp") //utk menambah data dipaling depan

	fmt.Println(data.Front().Next().Next().Value) //utk mengambil data sesuai urutan banyaknya next
	fmt.Println(data.Back().Prev().Prev().Value)

	fmt.Println(data.Front().Value) // utk mengambil data paling depan
	fmt.Println(data.Back().Value)  // utk mengambil data paling belakang

	//iterasi dari depan ke belakang
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	//iterasi dari belakang ke depan
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}
