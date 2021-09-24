package main

import "fmt"

func main() {

	name := "Bayu"

	if name == "Bayu" {
		fmt.Println("Hello Bayu")
	} else if name == "Bagus" {
		fmt.Println("Hello Bagus")
	} else if name == "Bagaswara" {
		fmt.Println("Hello Bagaswara")
	} else {
		fmt.Println("Hi, Boleh Kenalan?")
	}

	// artinya kita cek dengan short statement, dimana kita cek panjang dari name
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
