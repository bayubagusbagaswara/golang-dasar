package main

import "fmt"

// representasi dari tipe data yang sudah ada

func main() {

	// artinya kita membuat tipe data baru bernama NoKTP
	// dan sebenarnya adalah alias dari string
	type NoKTP string
	type Married bool

	var ktpBayu NoKTP = "35711111"
	var marriedStatus Married = true

	fmt.Println(ktpBayu)
	fmt.Println(marriedStatus)
}
