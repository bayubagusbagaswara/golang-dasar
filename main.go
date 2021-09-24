package main

import "fmt"

func main() {

	// contoh break
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Perulangan ke-", i)
	}
	// saat i = 5, maka langsung break keluar perulangan, sudah tidak menjalankan perulangan lagi
	// jadi hanya perulangan index ke 0, 1, 2, 3, 4

	fmt.Println("=========")

	// contoh continue (skipping)
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Perulangan ke-", i)
	}
	// saat i adalah genap maka perulangan akan diskip dan lanjut ke perulangan index selanjutnya (tidak keluar dari perulangan sampai kondisi perulangan bernilai true)
	// jadi hanya akan mengeksekusii jika perulangan berindex ganjil

}
