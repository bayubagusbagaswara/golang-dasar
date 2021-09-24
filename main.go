package main

import "fmt"

func main() {

	// contoh membuat constant
	// juga tidak wajib menuliskan jenis tipe datanya
	// Bedanya dengan varible, meskipun constant tidak pernah digunakan dia tidak akan complain atau error
	// Bisa juga dibuat mulitple deklarasi

	const (
		firstName string = "Bayu"
		lastName         = "Bagaswara"
		value            = 1000
	)

	// error
	// firstName = "Tidak bisa diubah"
	// lastName = "Tidak bisa diubah"

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

}
