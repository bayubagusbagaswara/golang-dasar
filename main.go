package main

import "fmt"

func main() {
	fmt.Println("Bayu")

	// menghitung karakter
	fmt.Println(len("Bayu Bagus"))

	// ambil karakter ke-0, hasilnya 10
	// karena 10 itu representasi dari byte nya karakter B
	fmt.Println("Bagaswara"[0])

	// mengambil karakter di index 1
	fmt.Println("Bayu Bagus Bagaswara"[1])
}
