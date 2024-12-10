package bench

import (
	"testing"
)

func BenchmarkInsideWithDummyPanic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fWithDummyPanic()
	}
}

func fWithDummyPanic() {
	defer func() {
		if r := recover(); r != nil {
		}
	}()
	f1WithDummyPanic()
}

func f1WithDummyPanic() (string, error) {
	// Генерация паники
	panic("Test panic")
}
