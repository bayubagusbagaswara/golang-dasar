package main

// saat ngelakuin import, maka secara otomatis function init() akan dieksekusi
// tanda _ agar saat kita sudah import package, tapi function didalam package tersebut tidak gunakan, hal tersebut akan menyebabkan program error
import (
	_ "golang-dasar/database"
)

func main() {

	// result := database.GetDatabase()

	// padahal function init() belum kita panggil (eksekusi), tapi sudah menampilkan datanya
	// fmt.Println(result)

}
