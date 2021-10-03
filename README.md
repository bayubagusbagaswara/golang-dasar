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
