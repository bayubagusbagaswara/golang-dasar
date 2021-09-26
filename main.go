package main

import "fmt"

func main() {

	counter := 0
	// didalam func bisa mengakses data dari luar
	// misal counter, bisa mengakses counter yang dideklarasikan di luar function increment
	increment := func() {
		fmt.Println("Increment")
		// hati-hati jangan sampe bikin function didalam function dan dapat mengubah data di function paling luarnya
		// seperti counter++ ini dapat mengubah nilai counter di function paling luarnya
		counter++
	}

	increment()          // Increment 1
	increment()          // Increment 2
	fmt.Println(counter) // 2
}
