package main

import (
	"fmt"
	"os"
)

func main() {

	// Function Args
	// misal ingin mengambil data argument dari sistem operasi kita, seperti lokasi file kita
	args := os.Args
	fmt.Println("Argument : ")
	fmt.Println(args)

	// Function Hostname
	// untuk mengambil nama Host dari sitem operasi kita
	// hostname balikannya ada 2 data
	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error", err.Error())
	}

	// Function Getenv
	// untuk mengambil environment dari sistem operasi kita
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	fmt.Println(username)
	fmt.Println(password)

}
