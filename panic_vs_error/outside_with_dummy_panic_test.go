package bench

import (
	"testing"
	"time"
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
	now := time.Now().UnixNano()
	if now < 1000 {
		panic("Test panic")
	} else {
		return "", nil
	}

}
