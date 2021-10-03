package main

import "fmt"

// misal kita punya struct bernama Address
type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Kediri", "Jawa Timur", "Indonesia"}

	// kita ingin address2 ini pointer ke address1, maka menggunakan operator &
	var address2 *Address = &address1

	// lalu kita coba mengubah data di address2
	address2.City = "Surabaya"

	fmt.Println(address1) // ini akan ikut berubah
	fmt.Println(address2)

}
