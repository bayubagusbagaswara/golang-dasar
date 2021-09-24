package main

import "fmt"

func main() {

	var a = 50
	var b = 10

	var jumlah = a + b
	var kali = a * b
	var kurang = a - b
	var bagi = a / b

	fmt.Println(jumlah)
	fmt.Println(kali)
	fmt.Println(kurang)
	fmt.Println(bagi)

	// augmented assignments
	var i = 10
	i += 10 // i = i + 10
	fmt.Println(i)

	// unary operator
	var j = 10
	j++ // j = j + 1
	fmt.Println(j)
	j-- // j = j - 1
	fmt.Println(j)

}
