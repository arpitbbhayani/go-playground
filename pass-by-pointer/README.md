Benchmarking Pass by Value and Pointer in Golang

## How to run

```
$ go test -gcflags=-N -bench=. -count 1
```

Sample output

```
goos: linux
goarch: amd64
pkg: github.com/arpitbbhayani/go-pointers-benchmark
cpu: AMD Ryzen 7 4800U with Radeon Graphics         
BenchmarkPBP-16         835952650                1.437 ns/op
BenchmarkPBV-16          1000000              1026 ns/op
PASS
ok      github.com/arpitbbhayani/go-pointers-benchmark  2.398s
```

## Generating Assembly Code

```
$ go tool compile -N -S ptr.go
```
