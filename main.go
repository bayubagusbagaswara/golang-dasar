package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(math.Ceil(1.40))  // dibulatkan keatas
	fmt.Println(math.Floor(1.60)) // dibulatkan kebawah
	fmt.Println(math.Round(1.60)) // dibulatkan yang paling dekat
	fmt.Println(math.Max(1, 2))   // membandingkan nilai yang paling besar
	fmt.Println(math.Min(1, 2))   // membandingkan nilai yang paling kecil
}
