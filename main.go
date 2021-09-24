package main

import "fmt"

func main() {
	// function main adalah function utama yang akan dieksekusi
	sayHello()

	for i := 0; i < 10; i++ {
		sayHello()
	}

}

func sayHello() {
	fmt.Println("Hello")
}
