package main

import (
	"strings"
)

// К каким негативным последствиям может привести данный фрагмент кода, и как
// это исправить? Приведите корректный пример реализации.

// var justString string

// func someFunc() {
// 	v := createHugeString(1 << 10)
// 	justString = v[:100]
// }

// func main() {
// 	someFunc()
// }

// Слайс фрагмента строки сохраняет указатель на оригинал строки, тем самым
// не давая сборщику мусора освободить память. Нужно явно копировать фрагмент,
// преобразовывая его в []byte.
// https://stackoverflow.com/questions/52395730/does-slice-of-string-perform-copy-of-underlying-data
// https://research.swtch.com/godata

var justString string

func createHugeString(count int) string {
	return strings.Repeat("a", count)
}

func someFunc() string {
	v := createHugeString(1 << 10)
	return string([]byte(v[:100]))
}

func main() {
	justString = someFunc()
}
