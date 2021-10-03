package main

import "fmt"

// struct tidak punya behavior(method didalamnya), hanya menyimpan data (field/properti/atribut)
type Customer struct {
	Name, Address string
	Age           int
}

// kita bikin function yang menempel ke struct Customer
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)

}
func main() {

	var bayu Customer
	bayu.Name = "Bayu"
	bayu.Address = "Indonesia"
	bayu.Age = 25

	// biasanya manggil functionnya langsung sebut nama function
	// sayHello(bayu, "Bayu")
	// karena function sayHello ditempel struct, maka sekarang manggilnya
	bayu.sayHello("Bayu")

}
