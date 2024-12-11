package recursion_bench

import (
	"errors"
	"flag"
	"os"
	"testing"
)

var errTest = errors.New("123")

//go:noinline
func ef(i, N int) error {
	if i < 0 {
		return errTest
	} else if i == N {
		return nil
	}
	return ef(i+1, N)
}

//go:noinline
func pf(i, N int) {
	if i < 0 {
		panic("test panic")
	} else if i == N {
		return
	}
	pf(i+1, N)
}

func testE(N int) {
	ef(0, N)
}

func testP(N int) {
	pf(0, N)
}

var deep int

func init() {
	flag.IntVar(&deep, "deep", 0, "Deep")
}

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

func BenchmarkPanic(b *testing.B) {

	for i := 0; i < b.N; i++ {
		testP(deep)
	}
}
func BenchmarkError(b *testing.B) {

	for i := 0; i < b.N; i++ {
		testE(deep)
	}
}