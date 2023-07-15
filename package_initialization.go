package main

import (
	"fmt"
	"test123/database"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
