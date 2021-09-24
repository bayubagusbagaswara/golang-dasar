# Tipe Data Slice

- Tipe data Slice adalah potongan dari data Array
- Slice mirip dengan Array, yang membedakan adalah ukuran Slice bisa berubah
- Slice dan Array selalu terkoneksi, dimana slice adalah data yang mengakses sebagian atau seluruh data di Array

## Detail Tipe Data Slice

- Memiliki 3 data yaitu pointer, length dan capacity
- Pointer adalah penunjuk data pertama di array pada slice
- Length adalah panjang dari slice
- Capacity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capacity

## Membuat Slice Dari Array

- array[low:high] , membuat slice dari array dimulai index low sampai index sebelum high
- array[low:] , membuat slice dari array dimulai index low sampai index akhir di array
- array[:high] , membuat slice dari array dimulai index 0 sampai index sebelum high
- array[:] , membuat slice dari array dimulai index 0 sampai index akhir di array

## Function Slice

- len(slice) , untuk mendapatkan panjang
- cap(slice) , untuk mendapatkan kapasitas
- append(slice, data) , membuat slice baru dengan menambah data ke posisi terakhir slice, jika kapasitas sudah penuh, maka akan membuat array baru
- make([]TypeData, length, capacity) , membuat slice baru
- copy(destination, source) , menyalin slice dari surce ke destination

## Hati-Hati Saat Membuat Array

- Jika salah, maka yang kita buat bukanlah, Array melainkan Slice
