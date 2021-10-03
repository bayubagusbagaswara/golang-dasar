package main

// membuat error dengan menggunakan package errors
import (
	"errors"
	"fmt"
)

// misal kit abikin function Pembagian
// return valuenya ada integer dan error
// error adalah interface, jadi bisa diset nil (kalau tidak terjadi error)
func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagi tidak boleh nol")
	} else {
		// error diganti dengan nil, karena pembagiannya berhasil/valid
		return nilai / pembagi, nil
	}
}

func main() {
	// panggil function Pembagian
	hasil, err := Pembagian(100, 20)
	// jika tidak terjadi error, maka print hasilnya
	if err == nil {
		fmt.Println("Hasil:", hasil)
	} else {
		// jika ada error, maka kita print errornya apa
		fmt.Println("Error:", err.Error())
	}
}
