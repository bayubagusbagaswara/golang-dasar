package main

import "fmt"

// misal kita punya function, dimana balikann adalah interface kosong
func random() interface{} {
	return true
}

func main() {
	result := random()

	// kita ingin konvert interface kosong menjadi string
	// ini jika kita emang yakni bahwa balikan interface kosongnya emang string
	// resultString := result.(string)
	// fmt.Println(resultString)

	// jika salah menggunakan type assertions, maka akan terjadi panic error
	// ini akan terjadi error panic, karena string tidak bisa dikonvert menjadi integer
	// resultInt := result.(int) // panic
	// fmt.Println(resultInt)

	// kita konversi result dengan type assertions
	switch value := result.(type) {
	case string:
		fmt.Println("String", value, "is string")
	case int:
		fmt.Println("Int", value, "is int")
	default:
		fmt.Println("Unknown type")
	}

}
