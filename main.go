package main

import (
	"fmt"
	"reflect"
)

// misal kita punya struct Sample
type Sample struct {

	// StructTag adalah menambah tag pada struct, menggunkan backtick ``
	Name string `required:"true" max:"10"`
	Age  int
}

// kegunaan reflect salah satunya adalah validation, daripada cek manual menggunakan if else
// misal kita ingin validasi data struct, datanya berupa interface kosong
// true kalo valid, false jika tidak valid
func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			// jika data nya sama dengan string kosong (""), maka dia tidak valid
			return reflect.ValueOf(data).Field(i).Interface() != ""
		}
	}
	return true
}

// kita ingin membaca berapa field dan sebagainya menggunakan reflect
func main() {
	// kita buat data struct dulu
	sample := Sample{"Bayu", 25}

	// kita buat reflectnya dari struct sample
	sampleType := reflect.TypeOf(sample)

	// lalu kita ambil fieldnya ke-0
	structField0 := sampleType.Field(0)
	structField1 := sampleType.Field(1)
	// ambil tag yang ada di struct
	required := structField0.Tag.Get("required")
	max := structField0.Tag.Get("max")

	fmt.Println(sampleType.NumField()) // mengetahui jumlah field
	fmt.Println(structField0.Name)     // untuk mengetahui nama field ke-0
	fmt.Println(structField1.Name)     // untuk mengetahui nama field ke-1
	fmt.Println(required)              // cetak tag
	fmt.Println(max)                   // cetak tag

	// kita panggil function untuk validation
	// yang kita validasi adalah data struct sample
	fmt.Println(IsValid(sample))
}
