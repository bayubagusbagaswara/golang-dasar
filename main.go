package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You Are Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	// Cara 1 membuat anonymous function dan disimpan dalam variable
	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("Bayu", blacklist)

	// Cara 2 membuat anonymous function langsung di parameter sebuah function
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}
