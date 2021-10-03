package main

import "fmt"

// function yang return valuenya adalah interface kosong
func Ups(i int) interface{} {
	// karena interface, maka return valuenya bisa interger, boolean, string dll
	// return 1
	// return true
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func main() {

	// harus disimpan kedalam variabel yang bertipe data interface{}
	var data interface{} = Ups(2)
	// variable kosong disini sudah menjadi interface
	fmt.Println(data)
}
