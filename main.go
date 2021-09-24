package main

import "fmt"

// function getFullName akan mengembalikan 2 return value yang bertipe string
// Ingat, setiap data yang ditreturnkan harus ditulis tipe datanya pada function nya
// meskipun tipe datanya sama, wajib harus tetep dituliskan tipe datanya untuk masing-masing data yang direturnkan

func getFullName() (string, string, int) {
	return "Bayu", "Bagaswara", 25
}

func getAddress() (string, string, int) {
	return "Kediri", "Indonesia", 41
}

func main() {

	// simpan return data dari getFullName kedalam 3 variabel yang berbeda
	// secara otomatis masing-masing data yang direturnkan berjumlah 3 data akan langsung disimpan

	firstName, lastName, age := getFullName()

	// tanda _ , artinya kita ignore return data dari getAdrress pada parameter ke 3 nya
	// juga untuk menghindari dari error oleh golang
	city, country, _ := getAddress()

	fmt.Println(firstName, lastName)
	fmt.Println("Age: ", age)
	fmt.Println("City: ", city)
	fmt.Println("Country: ", country)
}
