package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strings.Contains("Bayu Bagaswara", "Bayu"))
	fmt.Println(strings.Split("Bayu Bagaswara", " "))
	fmt.Println(strings.ToLower("Bayu Bagaswara"))
	fmt.Println(strings.ToUpper("Bayu Bagaswara"))
	fmt.Println(strings.Trim("    Bayu Bagaswara     ", " "))
	fmt.Println(strings.ReplaceAll("Bayu Bayu Bayu Bayu", "Bayu", "Bagus"))

}
