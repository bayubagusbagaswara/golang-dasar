package main

import "fmt"

func main() {
	// cara menulis array, harus gandeng dengan tipe data yang diinginkan [100]integer
	var names [3]string

	// isi data ke arraynya melalu index
	names[0] = "Bayu"
	names[1] = "Bagus"
	names[2] = "Bagaswara"

	// cara ambil data didalam array juga harus melalui indexnya
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// membuat array langsung saat deklarasi
	var value = [3]int{
		12,
		19,
		96,
	}
	fmt.Println(value) // satu array yang dicetak

	// hitung panjang array, bukan jumlah data
	// karena bisa saja datanya belum ada, tapi arraynya sudah terinisialisasi panjang (kapasitas) arraynya
	fmt.Println(len(names))
	fmt.Println(len(value))

}
