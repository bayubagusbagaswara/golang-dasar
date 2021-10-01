package main

import "fmt"

// misal kita function logging()
func logging() {
	fmt.Println("Selesai memanggil function")
}

// function runApplication()
func runApplication(value int) {

	// defer artinya akan menjalankan logging() setelah function runApplication selesai dijalankan
	// meskipun di runApplication terjadi error, maka defer akan tetap dijalankan
	defer logging()
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println("Result: ", result)
}

// function endApp, untuk mencoba Panic
func endApp() {
	// ketika terjadi panic, kita tangkap dengan recover(), agar aplikasinya tidak berhenti
	message := recover()
	// jika message nil, maka tidak terjadi error panic
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}
	fmt.Println("Aplikasi Selesai")
}

func runApp(error bool) {
	// defer tetap dijalankan meskipun terjadi error
	defer endApp()
	// jika terjadi error, maka jalankan panic
	if error {
		panic("Aplikasi Error")
	}
	// ketika terjadi panic, maka kode program dibawahnya tidak akan dijalankan
	// oleh karena itu, untuk recover kita harus taruh di paling atas
	// salah satu cara, kita taruh didalam function endApp
	// karena endApp ini adalah defer, maka dia akan tetap dieksekusi
	fmt.Println("Aplikasi Berjalan")

}

func main() {

	// panggil function runApplication
	// dan misal kita kirimkan parameter yang menyebabkan function error
	// runApplication(0)
	runApplication(10)

	// panggil function runApp, dan aktifkan error
	runApp(false)
	// untuk memastikan aplikasinya tidak berhenti setelah direcover
	fmt.Println("Aplikasi tidak berhenti")

}
