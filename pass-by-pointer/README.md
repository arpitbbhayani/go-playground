Benchmarking Pass by Value and Pointer in Golang

## How to run

```
$ go test -gcflags=-N -bench=. -count 1
```

Sample output

```
goos: linux
goarch: amd64
pkg: github.com/arpitbbhayani/go-playground/pass-by-pointer
cpu: AMD Ryzen 9 7950X 16-Core Processor            
BenchmarkPBP-32         1000000000               0.7663 ns/op
BenchmarkPBV-32          1122776              1145 ns/op
PASS
ok      github.com/arpitbbhayani/go-playground/pass-by-pointer  2.750s
```

## Generating Assembly Code

```
$ go tool compile -N -S -l ptr.go
```
