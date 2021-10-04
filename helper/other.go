package helper

import "fmt"

// tapi kalau didalam satu package TIDAK BOLEH ADA NAMA FUNCTION YANG SAMA, HARUS BERBEDA
// function dengan nama SayHello tidak bisa dituliskan disini, karena sudah digunakan di file lain (dalam satu package yang sama)
// func SayHello(name string) {

// }

// nama function harus berbeda
func Other(name string) {
	fmt.Println("Other", name)
}
