package main

import "fmt"

func main() {

	// deklarasi variable name bertipe data string
	var name string

	// assign variablenya
	name = "Bayu Bagus"
	// cetak isi variable name
	fmt.Println(name)

	// assign variable lagi
	name = "Aan Putra"
	// cetak isi variable name
	// hasilnya akan menimpa isi variable sebelumnya
	fmt.Println(name)

	// inisialisasi (mengisi nilai) variable tanpa harus menyebutkan tipe datanya
	var nama = "Bayu Bagus Bagaswara"
	fmt.Println(nama)

	nama = "Bayu Bagaswara"
	fmt.Println(nama)

	// bisa juga kita tambahkan tipe datanya, jika emang keharusan, misal harus int8
	var age int8 = 30
	fmt.Println(age)

	// tanpa menuliskan keyword var, diganti dengan :=
	// tapi hanya diawal deklarasi saja, selanjutnya bisa memakai = biasa

	country := "Indonesia"
	fmt.Println(country)

	// ganti isi variabelnya, cukup pake =
	country = "America"
	fmt.Println(country)

	// Deklarasi Multiple Variable
	var (
		firstName = "Bayu Bagus"
		lastName  = "Bagaswara"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

}
