# Interface Kosong

- Go-Lang bukanlah bahasa pemrograman yang berorientasi objek
- Biasanya dalam pemrograman berorientasi objek, ada satu data parent di puncak yang bisa dianggap sebagai semua implementasi data yang ada di bahasa pemrograman tersebut
- Contoh di Java ada java.lang.Object
- Untuk menangani kasus seperti ini, di Go-Lang kita bisa menggunakan interface kosong
- Interface kosong adalah interface yang `tidak memiliki deklarasi method satupun`. Hal ini membuat secara otomatis semua tipe data akan menjadi implementasinya

## Penggunaan Interface Kosong di Go-Lang

- Ada banyak contoh penggunaan interface kosong di Go-Lang, seperti:
  - fmt.Println(a ...interface{})
  - panic(v interface{})
  - recover() interfac{}
  - dan lain-lain
