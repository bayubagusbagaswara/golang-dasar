package main

import (
	"fmt"
	"regexp"
)

func main() {

	// kita buat regex, dimana karakter pertama `b` dan karakter akhir `u`
	// untuk karakter yang tengah bebas, antara a-z
	var regex = regexp.MustCompile(`b([a-z][a-z]u)`)

	fmt.Println(regex.MatchString("bayu"))
	fmt.Println(regex.MatchString("bedu"))
	fmt.Println(regex.MatchString("baku"))
	fmt.Println(regex.MatchString("biru"))

	// cari semua string berdasarkan sebuah regex
	fmt.Println(regex.FindAllString("bayu baku bati b1u bwu", 10))

}
