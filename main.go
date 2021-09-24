package main

import (
	"fmt"
)

func main() {

	name := "Bayu"

	switch name {
	case "Bayu":
		fmt.Println("Hello Bayu")
		fmt.Println("Hello bbb")
	case "Bagus":
		fmt.Println("Hello Bagus")
		fmt.Println("Hello cccc")
	default:
		fmt.Println("Hi, Boleh Kenalan?")
		fmt.Println("Hi, kenalan donk?")
	}

	// ini bukan kondisi, tapi pengecekan data, hasilnya dimasukkan ke case
	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("Nama terlalu panjang")
	// case false:
	// 	fmt.Println("Nama sudah benar")
	// }

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
