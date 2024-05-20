package main

import "testing"

func BenchmarkNewPersonPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewPersonPointer("Alice", 31)
	}
}

func BenchmarkNewPersonCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewPersonCopy("Alice", 31)
	}
}
