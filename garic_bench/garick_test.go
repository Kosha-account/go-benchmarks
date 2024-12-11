package garic_bench

import (
	"errors"
	"testing"
)

var errTest = errors.New("123")

//go:noinline
func ef(i, N int) error {
	if i == N {
		return errTest
	}

	return ef(i+1, N)
}
func testE(N int) {
	ef(0, N)
}

//go:noinline
func pf(i, N int) {
	if i == N {
		panic("123")
	}
	pf(i+1, N)
}

func testP(N int) {
	defer func() {
		r := recover()
		if r != nil {
			_ = r
		}
	}()
	pf(0, N)
}
func BenchmarkPanic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testP(b.N)
	}
}
func BenchmarkError(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testE(b.N)
	}
}
