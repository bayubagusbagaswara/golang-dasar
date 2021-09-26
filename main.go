package main

import "fmt"

// Factorial dengan Program Perulangan Biasa
func factorialLoop(value int) int {
	// buat variable result bernilai 1
	result := 1

	// buat perulangan untuk melakukan factorial
	// factorial akan terus berkurang nilainya setelah ia dikalikan
	// result = result * i
	// 5x4x3x2x1
	for i := value; i > 0; i-- {
		result *= i // result = result * i
		fmt.Println("Result: ", result)
	}
	return result
}

// Factorial dengan Recursive
func factorialRecursive(value int) int {

	// value 1 berfungsi untuk menghentikan function recursive ini sendiri
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1) // value * (value-1)
	}
}

func main() {
	loop := factorialLoop(5)
	fmt.Println(loop)

	recursive := factorialRecursive(5)
	fmt.Println(recursive)
}
