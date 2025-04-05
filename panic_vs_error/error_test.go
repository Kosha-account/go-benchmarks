package bench

import (
	"errors"
	"fmt"
	"log"
	"testing"
	"time"
)

func BenchmarkError(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fError0()
	}
}

func fError0() {
	_, err := fError1()
	if err == nil {
		log.Println(err)
	}
}

func fError1() (string, error) {
	now := time.Now().UnixNano()
	if now > 1000 {
		return "", errors.New("stop")
	} else {
		return fmt.Sprint("no error", now), nil
	}
}
