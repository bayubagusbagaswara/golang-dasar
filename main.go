package main

import (
	"flag"
	"fmt"
)

func main() {

	// defaultnya adalah localhost, root, root
	host := flag.String("host", "localhost", "Put your database host")
	username := flag.String("username", "root", "Put your database username")
	password := flag.String("password", "root", "Put your database password")

	// untuk melakukan proses parsing
	flag.Parse()

	// setelah diparse baru bisa kita menggunakan variable-variable tersebut
	// arahkan ke pointer host, username, password
	fmt.Println("Host : ", *host)
	fmt.Println("User : ", *username)
	fmt.Println("Password : ", *password)

}
