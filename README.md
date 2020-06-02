# go_powtest
Short program to benchmark square and 4th powers of float64 using multiplication vs math.Pow():

Results on a  i7-7700K is overclocked to 5GHz in single-threaded loads:

```
>go test -bench .
goos: windows
goarch: amd64
pkg: powtest
BenchmarkSquareByMultiplying-8          2000000000               0.43 ns/op
BenchmarkSquareWithPow-8                100000000               21.4 ns/op
BenchmarkQuadByMultiplying-8            2000000000               0.21 ns/op
BenchmarkQuadWithPow-8                  50000000                22.0 ns/op
```
