package main

import "fmt"

func main() {

	counter := 1

	// counter ke-11 sudah tidak dieksekusi perulangannya
	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++ // increment
	}

	// for dengan statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	slice := []string{"Bayu", "Bagus", "Bagaswara"}

	// untuk ambil data dari slice bisa menggunakan perulangan
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// cara cepat mengambil data adalah menggunakan for range
	names := []string{"Bayu", "Bagus", "Bagaswara"}
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}

	// atau tanpa index
	for _, name := range names {
		fmt.Println(name)
	}

	person := make(map[string]string)
	person["name"] = "Bayu"
	person["title"] = "Programmer"

	// for range juga bisa digunakan pada MAP
	for key, value := range person {
		fmt.Println(key, "=", value)
	}

}
