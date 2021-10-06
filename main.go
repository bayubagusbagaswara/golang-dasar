package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// misal kita punya function
func ChangeAddressToIndonesia(address *Address) {
	// kita coba mengubah data address
	// tidak bisa mengubah data address
	// jika ingin data address berubah, maka parameter function ini harus pointer

	address.Country = "Indonesia"
}

func main() {

	// contoh parameter function yang bukan Pointer
	address := Address{"Kediri", "Jawa Timur", ""}
	// panggil function dan kirimkan parameter berupa struct
	// yang dikirim ke function ini adalah valuenya (bukan referencenya)
	// kita sudah ubah parameternya adalah pointer, oleh karena itu, yang kita kirim adalah pass by reference (tambahkan tanda &)
	ChangeAddressToIndonesia(&address)

	fmt.Println(address) // data asli address tidak berubah

	// Note:
	// jika membuat Data Struct yang besar, maka lebih baik setting parameternya berupa pointer
	// perubahan data akan disimpan dalam memori, semakin banyak data maka memori penuh
	// jika memori penuh, aplikasi akan lambat

}
