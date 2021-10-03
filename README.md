# Pointer

## Pass by Value

- Secara default di Go-Lang semua variable itu di passing by value, bukan by reference
- Artinya, jika kita mengirim sebuah variable ke dalam function, method atau variable lain, sebenarnya yang dikirim adalah duplikasi value nya

## Pointer

- Pointer adalah kemampuan membuat reference ke lokasi data di `memory yang sama`, tanpa menduplikasi data yang sudah ada
- Sederhananya, dengan kemampuan pointer, kita bisa membuat `pass by reference`
- Pass by Reference itu kebalikan dari Pass by Value
- Kalau pass by reference datanya akan tetap sama dilokasi yang sama pula (tidak diduplikat)

## Operator &

- Untuk membuat sebuah variable dengan nilai pointer ke variable yang lain, kita bisa menggunakan operator `&` diikuti dengan nama variable nya

## Operator `*`

- Saat kita mengubah variable pointer, maka data yang berubah hanya variable tersebut
- Semua variable yang mengacu ke data yang sama tidak akan berubah
- Jika kita ingin mengubah seluruh variable yang mengacu ke data tersebut, kita bisa menggunakan operator `*`

## Function new

- Sebelumnya untuk membuat pointer dengan menggunakan operator `&`
- Go-Lang juga memiliki function new yang bisa digunakan untuk membuat pointer
- Namun, function `new` hanya mengembalikan pointer ke data kosong, artinya tidak ada data awal
