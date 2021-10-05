package main

import (
	"fmt"
	"time"
)

func main() {

	// membuat waktu saat ini
	now := time.Now()
	// bisa ambil data detik, jam, tanggal dll
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	// membuat waktu secara manual
	utc := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(utc.Local())

	// jika ingin parsing waktu, atau mengubah format penulisan waktu
	// layout := "yyyy-MM-dd" // kalo di Java
	layout := time.RFC3339 // kalo di Golang
	parse, _ := time.Parse(layout, "2006-01-02T15:04:05Z")
	fmt.Println(parse)

}
