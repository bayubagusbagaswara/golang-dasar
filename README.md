# Tipe Data Map

- Pada Array atau Slice, untuk mengakses data kita menggunakan index Number dimulai dari 0
- Map adalah tipe data lain yang berisikan kumpulan data yang sama. Namun, kita bisa menentukan jenis tipe data index yang akan kita gunakan
- Sederhananya, Map adalah tipe data kumpulan `key-value` (kata kunci - nilai), dimana kata kuncinya bersifat unik, tidak boleh sama
- Berbeda dengan Array dan Slice, jumlah data yang kita masukkan ke dalam Map boleh sebanyak-banyaknya, asalkan kata kuncinya (key) berbeda. Jika kita menggunakan kata kunci sama, maka secara otomatis data sebelumnya akan diganti dengan data baru

## Function Map

- len(map) , untuk mendapatkan jumlah data di map
- map[key] , mengambil data di map dengan key
- map[key] = value , mengubah data (value) di map dengan key
- make(map[TypeKey]TypeValue) , membuat map baru
- delete(map, key) , menghapus data di map dengan key
