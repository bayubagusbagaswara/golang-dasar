package main

import "fmt"

func main() {

	// map[typeKey][typeValue]
	person := map[string]string{
		"name":    "Bayu",
		"address": "Kediri",
	}

	// kalo di array jika kita ingin akses datanya menggunakan index
	// tapi di map untuk akses data di map mengguakan key nya

	fmt.Println(person)            // cetak mapnya
	fmt.Println(person["name"])    // ambil isi map dengan key "name"
	fmt.Println(person["address"]) // ambil isi map dengan key "address"

	// menambah atau memanipulasi data ke dalam map
	person["title"] = "Programmer"
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Buku Go-Lang"
	book["author"] = "Bayu Bagaswara"
	book["wrong"] = "Ups"

	fmt.Println(book) // sebelum dihapus
	fmt.Println(len(book))

	delete(book, "wrong") // hapus data di map book dengan key "wrong"

	fmt.Println(book) // setelah dihapus
	fmt.Println(len(book))

}
