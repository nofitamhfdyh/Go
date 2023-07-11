package main

import (
	"fmt"
)

// (1)
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i // result = result * i
	}

	return result
}

// (2)
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	//recursive function = func yang memanggil dirinya sendiri
	// cth kasusnya utk menyelesaikan factorial
	// caranya menggunakan looping(1) dan recursive(2)

	//(1)
	loop := factorialLoop(5)
	fmt.Println(loop)
	// fmt.Println(5 * 4 * 3 * 2 * 1)

	//(2)
	fmt.Println(factorialRecursive(5))
}
