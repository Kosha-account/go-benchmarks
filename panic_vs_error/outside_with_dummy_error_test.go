package bench

import (
	"errors"
	"fmt"
	"log"
	"testing"
	"time"
)

func BenchmarkOutsideWithDummyError(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fOutsideWithDummyError()
	}
}

func fOutsideWithDummyError() {
	_, err := f1OutsideWithDummyError()
	if err != nil {
		log.Println(err)
	}
}

func f1OutsideWithDummyError() (string, error) {
	now := time.Now().UnixNano()
	if now < 1000 {
		return "", errors.New(fmt.Sprintf("%d", now))
	} else {
		return "", nil
	}

}
