package main

import "fmt"

// misal kita bikin kontrak HasName
// hanya berisi kontrak method
type HasName interface {
	GetName() string
}

// interface juga bisa digunakan sebagai parameter di sebuah function
// misal function sayHello ini parameternya adalah berupa interface
// didalam function ini bisa memanggil kontrak method milik interface
func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

// misal kita punya struct Person
type Person struct {
	Name string
}

// misal kita buat struct method dimana namanya adalah GetName()
// karena dia punya function GetName(), maka secara otomatis struct Person akan mengikuti kontrak interface HasName yang memiliki function GetName()
func (person Person) GetName() string {
	return person.Name
}

// struct Animal
type Animal struct {
	Name string
}

// struct function Animal dengan nama GetName() dan return valuenya string
func (animal Animal) GetName() string {
	return animal.Name
}

func main() {

	// kalau kita manggil function SayHello kita gak bisa langsung memasukkan parameternya, karena parameternya berupa interface
	// caranya adalah kita harus membuat implementasi dari interface tersebut
	person := Person{Name: "Bayu"}
	// panggil function sayHello
	SayHello(person)

	cat := Animal{Name: "Push"}
	SayHello(cat)
}
