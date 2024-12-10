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
BenchmarkInsideWithDummyError-16        43333726                26.01 ns/op
BenchmarkInsideWithDummyPanic-16        20398630                55.97 ns/op
BenchmarkInsideWithError-16             1000000000               0.1375 ns/op
BenchmarkInsideWithPanic-16             20759938                57.01 ns/op
BenchmarkOutsideWithDummyError-16       45638512                25.87 ns/op
BenchmarkOutsideWithDummyPanic-16       1000000000               0.0000011 ns/op
PASS
ok      go-benchmarks/panic_vs_error    4.968s

```
