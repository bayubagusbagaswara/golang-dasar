package main

import "fmt"

type Man struct {
	Name string
}

// kita sudah tambahkan pointer di method struct ini
func (man *Man) Married() {
	// harapannya kita akan ubah Name nya ditambahi Mr
	// tapi yang di passing disini adalah Valuenya (duplikasi valuenya), dan reference tidak akan berubah
	// didalam method ini Name nya sudah berubah
	// tapi begitu keluar method, maka nilai name akan tetap sama yang asli (tidak berubah)
	// untuk mengatasinya, maka tambahkan Pointer di method
	// dengan pointer akan lebih hemat memory, karena tidak menduplikasi banyak data
	man.Name = "Mr. " + man.Name
	fmt.Println("Name di method ini:", man.Name)
}

func main() {

	bayu := Man{"Bayu"}
	bayu.Married()

	fmt.Println(bayu.Name) // hasilnya Bayu, karena data aslinya tidak berubah

}
