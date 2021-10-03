package main

import "fmt"

// diawali keyword type, lalu diikuti nama struct nya, dan tambahkan keyword struct
type Customer struct {
	Name, Address string
	Age           int
}

func main() {

	// buat variable bertipe data struct
	var bayu Customer

	// mengubah atau mengassign data dari variable bertipe data struct
	bayu.Name = "Bayu Bagus Bagaswara"
	bayu.Address = "Indonesia"
	bayu.Age = 25

	fmt.Println(bayu)
	fmt.Println(bayu.Name)
	fmt.Println(bayu.Address)
	fmt.Println(bayu.Age)

	// membuat data struct secara langsung
	aan := Customer{
		Name:    "Aan Putra",
		Address: "Kediri",
		Age:     25,
	}

	// ini urutannya harus sesuai
	agung := Customer{"Agung", "Jakarta", 25}

	fmt.Println(aan)
	fmt.Println(agung)
}
