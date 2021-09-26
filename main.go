package main

import "fmt"

// buat function getGoodBye
func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {

	// isi variabel goodbye dengan function lain (tapi hanya nama functionnya)
	// jadi function getGoodBye disini adalah sebagai value
	goodbye := getGoodBye

	// lalu kita bsia memasukkan nilai parameter function nya di variable yang menyimpan function utamanya
	result := goodbye("Bayu")
	fmt.Println(result)

	// hasilnya akan sama saat kita langsung memanggil function getGoodBye
	// tapi pada cara pertama kita bisa memasukkannya ke dalam sebuah variable
	fmt.Println(getGoodBye("Bayu"))

}
