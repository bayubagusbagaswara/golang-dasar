package database

import "fmt"

var connection string

// function init() ini akan otomatis dieksekusi saat package database ini diakses dari luar
func init() {
	fmt.Println("Function Init dipanggil")
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
