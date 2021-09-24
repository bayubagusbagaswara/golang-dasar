package main

import "fmt"

func main() {

	var name1 = "Bayu"
	var name2 = "Bayu"

	// membandingkan dua nilai dari name1 dan name2
	var result bool = name1 == name2

	fmt.Println(result)

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
