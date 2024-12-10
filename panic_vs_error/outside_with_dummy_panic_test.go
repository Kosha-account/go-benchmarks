package bench

import (
	"testing"
)

func BenchmarkOutsideWithDummyPanic(b *testing.B) {
	defer func() {
		if r := recover(); r != nil {
		}
	}()
	for i := 0; i < b.N; i++ {
		fOutsideWithDummyPanic()
	}
}

func fOutsideWithDummyPanic() {
	f1OutsideWithDummyPanic()
}

func f1OutsideWithDummyPanic() (string, error) {
	// Генерация паники
	panic("Test panic")
}
