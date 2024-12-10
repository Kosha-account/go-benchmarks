package bench

import (
	"errors"
	"testing"
)

func BenchmarkInsideWithError(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fWithError()
	}
}

func fWithError() {
	_, err := f1WithError()
	if err != nil {
	}
}

func f1WithError() (string, error) {
	return "", errors.New("Test error")
}
