# Defer

- `Defer function` adalah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai di eksekusi
- Kadang kita butuh mengeksekusi sesuatu setelah sebuah function itu dieksekusi
- Defer function akan selalu dieksekusi walaupun terjadi `error` di function yang dieksekusi

# Panic

- Panic function adalah function yang bisa kita gunakan untuk menghentikan program
- Panic function biasanya dipanggil ketika terjadi error pada saat program kita berjalan
- Saat panic function dipanggil, program akan terhenti. Namun, `defer function` tetap akan dieksekusi

# Recover

- Recover adalah function yang bisa kita gunakan untuk menangkap data panic
- Dengan recover proses panic akan terhenti, sehingga program akan `tetap berjalan`
- Ketika terjadi `panic` terus kita `recover`, maka aplikasinya tidak jadi berhenti
