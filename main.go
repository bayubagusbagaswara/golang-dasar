package main

import "fmt"

// buat function getFullName
// kita bisa tambahkan tipe data pada named parameternya, secara otomatis itu akan direturnkan dari function ini
func getFullName() (firstName string, middleName string, lastName string) {

	// kita assign
	firstName = "Bayu"
	middleName = "Bagus"
	lastName = "Bagaswara"

	// cukup tulis return, maka secara otomatis semua nilai akan direturnkan
	// ini terjadi karena kita sudah mendeklarasikan variable return value di parameter named semua data yang ada di function ini
	return
}

func main() {
	// secara otomatis funciton getFullName mempunyai 3 buat return data
	// kita tampung data-data tersebut kedalam sebuah variable
	a, b, c := getFullName()

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
