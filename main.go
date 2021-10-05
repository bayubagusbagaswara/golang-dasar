package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {

	// buat data ring dengan jumlah data ring 5
	// jadi tidak bisa nambah data seenaknya
	data := ring.New(5)

	// set data kedalam ring
	for i := 0; i < data.Len(); i++ {
		// set menggunakan Value, dan isi valuenya apa
		data.Value = "Value " + strconv.FormatInt(int64(i), 10)
		// untuk pindah ke data ring selanjutnya menggunakan Next()
		data = data.Next()
	}

	// iterasi untuk ambil data ring menggunakan function Do()
	data.Do(func(value interface{}) {
		fmt.Println(value)
	})
}
