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
BenchmarkInsideWithDummyError-16        39657985                25.88 ns/op
BenchmarkInsideWithDummyPanic-16        21379144                57.59 ns/op
BenchmarkInsideWithError-16             1000000000               0.1217 ns/op
BenchmarkInsideWithPanic-16             19744482                59.06 ns/op
BenchmarkOutsideWithDummyError-16       45543954                26.09 ns/op
BenchmarkOutsideWithDummyPanic-16       46769871                25.69 ns/op
PASS
ok      go-benchmarks/panic_vs_error    6.161s

```
