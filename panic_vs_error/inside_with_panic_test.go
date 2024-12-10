package bench

import (
	"testing"
)

func BenchmarkInsideWithPanic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fWithPanic()
	}
}

func fWithPanic() {
	defer func() {
		if r := recover(); r != nil {
		}
	}()
	f1WithPanic()
}

func f1WithPanic() (string, error) {
	panic("Test panic")
}
