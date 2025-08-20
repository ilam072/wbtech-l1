package main

import "strings"

//var justString string

func main() {
	someFunc()
}

func someFunc() {
	v := createHugeString(1 << 10)

	// Проблема: justString сохраняет срез от v, и хотя используется только 100 символов,
	// под капотом justString ссылается на весь массив v.
	// Поскольку justString — глобальная переменная, сборщик мусора не сможет освободить память,
	// занятую всей строкой v (1024 байта).
	// justString = v[:100]

	// Решение:
	// Создаст копию в новой области памяти, и justString не будет ссылаться на v
	justString := strings.Clone(v[:100])

	_ = justString

}

func createHugeString(size int) string {
	return strings.Repeat("а", size)
}
