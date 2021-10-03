package main

import "fmt"

// kita bikin function NewMap
func NewMap(name string) map[string]string {
	if name == "" {
		// default value dari tipe data string adalah string kosong
		// sehingga kita bisa menggantinya ebagai nil
		return nil
	} else {
		// jika name nya tidak kosong, maka return data map nya
		return map[string]string{
			"name": name,
		}
	}
}

func main() {

	// jika ingin menginisialisasi data map sebagai nil
	var person map[string]string = nil
	fmt.Println(person)
	// pengecekan tanpa nil
	// biasanya kita mengecek melalui Key di map, apakah key tersebut memiliki data atau tidak
	if person["name"] == "" {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(person)
	}

	// pengecekan menggunakan nil
	// biasanya nil digunakan untuk mengecek data, apakah data tersebut ada atau tidak
	data := NewMap("")
	if data == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(data)
	}

}
