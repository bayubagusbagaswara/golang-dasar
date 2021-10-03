# Type Assertions

- Type Assertions merupakan kemampuan `mengubah tipe data` menjadi tipe data yang diinginkan
- Fitur ini sering sekali digunakan ketika kita `bertemu` dengan data `interface kosong`

## Type Assertions Menggunakan Switch

- Saat salah menggunakan type assertions, maka bisa berakibat terjadi panic di aplikasi kita
- Jika panic dan tidak ter-recover, maka otomatis program kita akan mati
- Agar lebih aman, sebaiknya kita menggunakan `switch expression` untuk melakukan type assertions
- Dengan switch ini, maka akan lebih aman. Karena panic akan di recover, dan aplikasi kita tidak akan mati
