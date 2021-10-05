package main

import (
	"fmt"
	"sort"
)

// misal kita punya struct
type User struct {
	Name string
	Age  int
}

// kita alias dari struct User, karena dia adalah Array of User
// User seperti Slice
// kita ingin UserSlice ini bisa diurutkan
type UserSlice []User

// kita buat method UserSlice
// kita implementasi Len() dari package sort
func (value UserSlice) Len() int {
	return len(value)
}

// implementasi Less()
// balikannya boolean
// jika i < j, maka balikan true
func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

// implementasi Swap(), value i akan diambil dari j, dan value j diambil dari value i
func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {

	users := []User{
		{"Bayu", 25},
		{"Bagus", 30},
		{"Bagaswara", 28},
		{"Aan", 19},
	}

	// kita akan melakukan pengurutan dari data users
	// caranya panggil package sort{}, lalu ada function Sort()
	// Sort() tidak ada return
	// jadi dia langsung mengurutkan
	sort.Sort(UserSlice(users))

	fmt.Println(users)

}
