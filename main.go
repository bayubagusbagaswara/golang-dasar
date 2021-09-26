package main

import "fmt"

// buat function sumAll
// artinya titik tiga (...) kita bisa menerima data lebih dari satu
// sehingga functon ini menjadi variadic
// dan data tersebut disimpan kedalam satu argument
// jika ingin menambahkan parameter lain juga bisa, misal parameter string
// tapi parameter varargs harus diletakkan paling belakang
func sumAll(numbers ...int) int {
	total := 0

	// numbers disini otomatis akan menjadi slice (kumpulan data)
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	// cukup memasukkan angka-angkanya
	// atau bahkan kita tidak wajib memasukkan datanya ke sumAll (tidak error)
	total := sumAll(10, 90, 30, 50, 40)
	fmt.Println(total)

	// menjadikan slice sebagi tipe parameter untuk Variadic Function
	// tapi tidak bisa langsung, harus disimpan dulu dalam variable
	// kita bikin variable slice numbers dulu
	sliceNumbers := []int{10, 10, 10, 10, 10}
	// masukkan numbers ke parameter di Variadic Function
	total = sumAll(sliceNumbers...)
	fmt.Println(total)
}
