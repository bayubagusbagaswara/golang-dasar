# Package sort

- Package sort adalah package yang berisikan utilitas untuk proses pengurutan
- Agar data kita bisa diurutkan, kita harus mengimplementasikan kontrak di interface `sort.Interface`

## Package sort.Interface

<!-- element of the collection be enumerated by an integer index -->

type Interface interface {

<!-- Len is the number of element in the collection -->

Len() int

<!-- Less reports whether the element with -->
<!-- index i should sort before the element with index j -->

Less(i, j int) bool

<!-- Swap swaps the elements with indexes i and j -->

Swap(i, j int)
}

- Len untuk mengembalikan data
- Less untuk cek lebih kecil dari
- Swap untuk menukar posisi
