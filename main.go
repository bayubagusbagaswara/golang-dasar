package main

import (
	"container/list"
	"fmt"
)

func main() {

	data := list.New()

	// untuk menambahkan data ke linked list menggunakan PushBack
	data.PushBack("Bayu")
	data.PushBack("Bagus")
	data.PushBack("Bagaswara")

	// ambil data linked list, caranya menggunakan iterasi satu persatu datanya
	// stop iterasi saat data sudah nil
	// Nil menandakan data linked sudah berada paling ujung
	// Front() adalah mengambil paling depan
	// Next() adalah data selanjutnya

	// contoh iterasi dari depan
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
	// contoh iterasi dari belakang
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
	fmt.Println("Data Paling Awal:", data.Front().Value)
	fmt.Println("Data Paling Akhir:", data.Back().Value)

}
