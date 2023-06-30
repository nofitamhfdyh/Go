package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
func main() {
	total := sumAll(10, 10, 10, 10)
	fmt.Println(total)

	//slice parameter
	slice := []int{90, 90, 90}
	total = sumAll(slice...)
	fmt.Println(total)
}
