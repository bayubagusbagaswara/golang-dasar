package main

import "fmt"

// balikan function adalah tipe data string
func getHello(name string) string {
	return "Hello " + name
	// kode program yang ditulis setelah kode program return, maka tidak akan dieksekusi
	// makanya return biasa ditulis paling akhir kode program
}

func main() {
	// balikan data dari getHello() akan disimpan kedalam variabel result
	result := getHello("Bayu")
	fmt.Println(result)
}
