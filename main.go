package main

import "fmt"

func main() {
	var ujian = 90
	var absensi = 80

	// cek data bool
	var lulusUjian bool = ujian > 80
	var lulusAbsensi bool = absensi > 80

	// cek lulus antara dua data bool
	var lulus bool = lulusUjian && lulusAbsensi

	fmt.Println(lulus) // false
	// karena absensi bernilai 80
}
