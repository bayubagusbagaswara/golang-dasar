# Package regexp

- Package regexp adalah utilitas di Go-Lang untuk melakukan pencarian regular expression
- Regular expression di Go-Lang menggunakan library C yang dibuat Google bernama RE2

## Beberapa Function di Package regexp

- regexp.MustCompile(string), untuk membuat Regexp
- Regexp.MatchString(string) bool, untuk mengecek apakah Regexp match dengan string. Misal digunakan cek format penulisan email dll
- Regexp.FindAllString(string, max), untuk mencari string yang match dengan maximum jumlah hasil
