package bench

import (
	"errors"
	"fmt"
	"log"
	"testing"
	"time"
)

func BenchmarkInsideWithDummyError(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fWithDummyError()
	}
}

func fWithDummyError() {
	defer func() {
		if r := recover(); r != nil {
		}
	}()
	_, err := f1WithDummyError()
	if err != nil {
		log.Println(err)
	}
}

func f1WithDummyError() (string, error) {
	now := time.Now().UnixNano()
	if now < 1000 {
		return "", errors.New(fmt.Sprintf("%d", now))
	} else {
		return "", nil
	}

}
