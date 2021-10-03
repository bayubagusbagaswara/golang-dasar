package main

import "fmt"

// misal kita punya struct bernama Address
type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Kediri", "Jawa Timur", "Indonesia"}

	// kita bikin variable address2 yang berasal dari address1
	address2 := address1

	// lalu kita coba mengubah data di address2
	address2.City = "Surabaya"

	fmt.Println(address1) // datanya tetap tidak berubah
	fmt.Println(address2)

}
