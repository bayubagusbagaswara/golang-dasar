package main

// kita ingin mengimport dari helper, maka tuliskan manual
// nama folder harus dimulai paling luar
// lebih enaknya kalau kita bikin go mod, dan melihat packagenya dari go mod
import "golang-dasar/helper"

func main() {

	helper.SayHello("Bayu")
	helper.Other("Aan")

}
