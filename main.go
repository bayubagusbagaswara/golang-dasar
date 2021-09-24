package main

import "fmt"

func main() {

	// bikin array months
	// panjang arraynya tidak ada batasan
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	// buat slice dari index-4 sampai index-7
	var slice1 = months[4:7]

	fmt.Println(slice1)      // [Mei Juni Juli], cetak array slice
	fmt.Println(len(slice1)) // 3, cetak panjang array slice
	fmt.Println(cap(slice1)) // 8, cetak kapasitas array slice

	// coba ubah data pada arraynya
	months[4] = "Diubah"
	fmt.Println(slice1) // hasilnya adalah slice1 juga ikut berubah

	// HATI-HATI ketika mengubah data didalam array
	// saat kita mengubah data di slicenya pun, data di array utama juga akan ikut berubah

	// misal kita ubah data di slice
	slice1[0] = "Mei lagi" // artinya index-0 di dalam array slice yang sebelumnya bernilai "Diubah"
	fmt.Println(slice1)
	fmt.Println(months) // di array Months pun datanya juga ikut berubah

	// Contoh Append Slice
	days := [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}

	// buat array slicenya, mulai dari index-5 pada array days, yakni Sabtu
	daySlice1 := days[5:] // [Sabtu Minggu]
	// ubah data array slice
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(days) // cetak array days, datanya juga akan berubah sesuai array slice

	// membuat array slice baru, yang isinya adalah array slice1 diappend (ditambah) data baru
	daySlice2 := append(daySlice1, "Libur Baru")
	// ubah data index-0 pada array slice2
	daySlice2[0] = "Ups"
	// cetak isi array slice2

	fmt.Println(daySlice2) // [Ups, Minggu Baru, Libur Baru]

	// Tapi ketika kita append, sedangkan kapasitas array utama (days) sudah full, maka data append tidak ditambahkan dan dibuatkan array baru
	fmt.Println(days) // [Senis, Selasa, Rabu, Kamis, Jumat, Sabtu Baru, Minggu Baru]

	// Ada cara lain, yakni bikin slice nya dari awal banget (tanpa ada deklarasi array sebelumnya)
	// buat slice bertipe String, lengthnya 2 dan kapasitasnya 5
	// sebenarnya akan dibuatkan array dengan panjang 5
	// ini lebih aman, karena array utamanya diumpetin dibalik slicenya
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Bayu"
	newSlice[1] = "Bagaswara"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// Program Copy Slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println(toSlice)

	// Perbedaan membuat Array dan Slice
	iniArray1 := [5]int{1, 2, 3, 4, 5}  // ini juga Array
	iniArray := [...]int{1, 2, 3, 4, 5} // harus ada titik ...
	iniSlice := []int{1, 2, 3, 4, 5}    // tanpa titik

	fmt.Println(iniArray)
	fmt.Println(iniArray1)
	fmt.Println(iniSlice)
}
