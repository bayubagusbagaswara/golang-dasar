package helper

var version = "1.0.0"      // tidak bisa diakses dari luar (package lain)
var Application = "golang" // bisa diakses dari luar

// function ini TIDAK BISA DIAKSES dari luar package
func sayGoodBye(name string) string {
	return "Good bye, " + name
}

// function ini BISA DIAKSES dari luar package
func SayHello(name string) string {
	return "Hello, " + name
}
