package main

import "fmt"

// misal kita punya struct bernama Address
type Address struct {
	City, Province, Country string
}

func main() {

	address1 := Address{"Kediri", "Jawa Timur", "Indonesia"}
	address2 := &address1 // address2 adalah pointer ke address1
	address3 := &address1

	// lalu kita coba mengubah data (field) di address2
	address2.City = "Surabaya"

	// gimana kalau kita mengubah keseluruhan variablenya?
	// caranya dengan menggunakan operator *
	// address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}

	// jika kita menambahkan tanda * di address2, maka dia akan memaksa address1 untuk pindah reference ke Address milik address2
	// jadi address1 sudah tidak reference ke Address{Kediri}
	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}

	fmt.Println(address1) // address1 berubah mengikuri address2
	fmt.Println(address2) // tapi address2 ini akan berpindah pointer ke Address yang baru
	fmt.Println(address3)

}
