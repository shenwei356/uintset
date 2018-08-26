[![GoDoc](https://godoc.org/github.com/shenwei356/uintset?status.svg)](https://godoc.org/github.com/shenwei356/uintset)
[![Go Report Card](https://goreportcard.com/badge/github.com/shenwei356/uintset)](https://goreportcard.com/report/github.com/shenwei356/uintset)

# uintset

    import "github.com/shenwei356/uintset"

Package uintset is a fast uint64 Set in golang, forked from https://github.com/brentp/intintmap


Performance

```
BenchmarkUintsetFill-4                                10         186584376 ns/op        201261312 B/op        24 allocs/op
BenchmarkIntIntMapFill-4                               5         226325438 ns/op        201195952 B/op        22 allocs/op
BenchmarkStdMapFill-4                                  5         308191604 ns/op        54742281 B/op      73511 allocs/op
BenchmarkUintsetTest10PercentHitRate-4             10000            113556 ns/op           20126 B/op          0 allocs/op
BenchmarkIntIntMapTest10PercentHitRate-4           10000            119272 ns/op           20119 B/op          0 allocs/op
BenchmarkStdMapTest10PercentHitRate-4              10000            128811 ns/op            5474 B/op          7 allocs/op
BenchmarkUintsetTest100PercentHitRate-4              500           2627686 ns/op          402522 B/op          0 allocs/op
BenchmarkIntIntMapTest100PercentHitRate-4            300           3969019 ns/op          670652 B/op          0 allocs/op
BenchmarkStdMapTest100PercentHitRate-4               100          13586000 ns/op          547218 B/op        734 allocs/op
```

## Usage

```go
m := intintmap.New(32768, 0.6)
m.Add(uint64(1234))
m.Add(uint64(123))

ok := m.Has(uint64(222))
ok := m.Has(uint64(333))

m.Del(uint64(222))
m.Del(uint64(333))

fmt.Println(m.Size())

for e := range m.Items() {
    fmt.Printf("element: %d,\n", e)
}

clone := m.Clone()
```
