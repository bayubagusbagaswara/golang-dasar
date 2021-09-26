package main

import (
	"fmt"
)

// type declarations digunakan untuk alias terhadap function
// sehingga saat menyebutkan function sebagai parameter, kita cukup panggil type declarations ini
type Filter func(string) string

// buat function sayHelloWithFilter
// parameter 1 adalah tipe data string
// parameter 2 adalah tipe data function
func sayHelloWithFilter(name string, filter Filter) {

	// buat variable yang menampung function filter
	nameFiltered := filter(name)

	// kita bisa memanggil parameter filter sebagai function
	// kita bisa memasukkan value parameter 1 kedalam function di parameter 2
	fmt.Println("Hello ", nameFiltered)
}

// misal kita buat function untuk memfilter kata-kata kasar
// jadi function yang bertugas memfilter terlebih dahulu sebelum mengeksekusi function yang utamanya
func spamFilter(name string) string {

	// jika name adalah anjing, maka akan kita ubah menjadi ...
	// jika bukan, maka kita return nilai asli name nya
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {

	// kita panggil function sayHelloWithFilter
	// masukkan parameter 1 berupa name string, parameter 2 berupa function spamFilter
	sayHelloWithFilter("Bayu", spamFilter) // Hello, Bayu

	// buat variable untuk menampung function spamFilter
	filter := spamFilter

	// panggil function sayHelloWithFilter
	// masukkan parameter string name, dan function filter
	sayHelloWithFilter("Anjing", filter) // Hello, ...
}
