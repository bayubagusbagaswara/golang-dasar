# Variadic Function

- Parameter yang berada di posisi terakhir pada sebuah function, itu memiliki kemampuan untuk dijadikan sebuah `varargs (variable argument)`
- `Variadic Function` artinya sebuah function yang memiliki variable argument
- `Varargs` artinya datanya bisa menerima `lebih dari satu input`, atau anggap saja semacam Array
- Apa bedanya dengan parameter biasa dengan tipe data Array?
  - Jika parameter tipe Array, kita wajib membuat array terlebih dahulu sebelum mengirimkan ke function
  - Jika parameter menggunakan `varargs`, kita bisa langsung mengirim datanya, jika lebih dari satu, maka cukup gunakan tanda koma `,`

## Slice Parameter

- Kadang ada kasus dimana kita menggunakan Variadic Function. Namun memiliki variable berupa `slice`
- Kita bisa menjadikan slice sebagai `varargs parameter`
