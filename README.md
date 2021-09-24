## Catatan

- Untuk menjalankan program harus di package main
- main function adalah tempat dimana aplikasi utama dijalankan
- Untuk menjalankan program golang, menggunakan cara compile satu persatu

## Urutannya

1. `go build <namafile.go>`
2. setelah dibuild, maka akan mendapatkan hasil binary `helloworld.exe`
3. compile binary `helloworld.exe` dengan cara `./helloworld`. Hasilnya adalah isi dari program yang telah dibuat yakni Hello World Colang!

## Tidak Efisien saat di Development

- Jika dengan cara compile satu persatu maka tidak akan efisien
- Ada cara lain yakni menggunakan `go run <namafile>`, tapi ini hanya dilakukan di development saja
- Jika di production harus tetap compile satu persatu (jangan langsung di run sekaligus), karena binary file hasil compile tersebut yang akan dibawa ke server

## Menghapus file binary

- menggunakan perintah `rm <nama file binarynya>`
