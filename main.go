package main

import "fmt"

func sayHello(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	// saat memanggil function sayHello maka wajib mengirimkan data sebagai parameter functionnya
	sayHello("Bayu", "Bagaswara")
}
