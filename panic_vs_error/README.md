# go-benchmarks

``` shell
cd panic_vs_error
go test -bench .
```

```
goos: linux
goarch: amd64
pkg: go-benchmarks/panic_vs_error
cpu: 12th Gen Intel(R) Core(TM) i7-12650H
BenchmarkInsideWithDummyError-16        43809744                26.21 ns/op
BenchmarkInsideWithDummyPanic-16        21024459                57.15 ns/op
BenchmarkInsideWithError-16             1000000000               0.1285 ns/op
BenchmarkInsideWithPanic-16             21038455                57.47 ns/op
BenchmarkOutsideWithDummyError-16       44460295                25.85 ns/op
BenchmarkOutsideWithDummyPanic-16       46023050                25.33 ns/op
PASS
ok      go-benchmarks/panic_vs_error    6.226s

```
