package bench

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func BenchmarkPanic(b *testing.B) {
	k := 0
	defer func() {
		if r := recover(); r != nil {
			//fmt.Println("finished with: +v", r)
			k++
		}
		fmt.Println("Just to not be skipped by compiler, k: +d", k, ", b.N: ", b.N)
	}()
	for i := 0; i < b.N; i++ {
		fPanic0()
	}
}

func fPanic0() {
	fPanic1()
}

func fPanic1() (string, error) {
	now := time.Now().UnixNano()
	if now > 1000 {
		panic(errors.New("test error"))
	} else {
		return fmt.Sprint("no error", now), nil
	}
}
