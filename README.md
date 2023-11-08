## levenshtein

go语言实现的Levenshtein Distance 算法及基于Levenshtein Distance 算法的相似度计算算法

## USAGE

```
package main

import (
	"fmt"

	"github.com/cloudfstrife/levenshtein"
)

func main() {
	fmt.Println(levenshtein.Levenshtein("你好世界", "晚安世界"))
	fmt.Println(levenshtein.SimilarDegree("你好世界", "晚安世界"))
}
```

输出：

```
$ go run main.go
2
0.5
```

## testing

```
$ go test -v -cover
=== RUN   TestLevenshtein
--- PASS: TestLevenshtein (0.00s)
=== RUN   TestSimilarDegree
--- PASS: TestSimilarDegree (0.00s)
PASS
coverage: 100.0% of statements
ok      github.com/cloudfstrife/levenshtein     0.043s

```

## benchmark

```
$ go test -v -bench="." -run="^$" -benchmem -cpu 2,4,6,8 github.com/cloudfstrife/levenshtein
goos: windows
goarch: amd64
pkg: github.com/cloudfstrife/levenshtein
BenchmarkLevenshtein/Second_Word_include_First_Word_-_2-2               10000000               171 ns/op              48 B/op          1 allocs/op
BenchmarkLevenshtein/Second_Word_include_First_Word_-_2-4               10000000               163 ns/op              48 B/op          1 allocs/op
BenchmarkLevenshtein/Second_Word_include_First_Word_-_2-6               10000000               161 ns/op              48 B/op          1 allocs/op
BenchmarkLevenshtein/Second_Word_include_First_Word_-_2-8               10000000               195 ns/op              48 B/op          1 allocs/op
BenchmarkLevenshtein/UNEqual-2                                           1000000              1293 ns/op             112 B/op          1 allocs/op
BenchmarkLevenshtein/UNEqual-4                                           1000000              1346 ns/op             112 B/op          1 allocs/op
BenchmarkLevenshtein/UNEqual-6                                           1000000              1226 ns/op             112 B/op          1 allocs/op
BenchmarkLevenshtein/UNEqual-8                                           1000000              1236 ns/op             112 B/op          1 allocs/op
BenchmarkLevenshtein/First_Word_include_Second_Word_-_1-2               10000000               193 ns/op              32 B/op          1 allocs/op
BenchmarkLevenshtein/First_Word_include_Second_Word_-_1-4               10000000               188 ns/op              32 B/op          1 allocs/op
BenchmarkLevenshtein/First_Word_include_Second_Word_-_1-6               10000000               186 ns/op              32 B/op          1 allocs/op
BenchmarkLevenshtein/First_Word_include_Second_Word_-_1-8               10000000               187 ns/op              32 B/op          1 allocs/op
BenchmarkLevenshtein/Second_Word_include_First_Word_-_1-2               10000000               198 ns/op              48 B/op          1 allocs/op
BenchmarkLevenshtein/Second_Word_include_First_Word_-_1-4               10000000               188 ns/op              48 B/op          1 allocs/op
BenchmarkLevenshtein/Second_Word_include_First_Word_-_1-6               10000000               176 ns/op              48 B/op          1 allocs/op
BenchmarkLevenshtein/Second_Word_include_First_Word_-_1-8               10000000               175 ns/op              48 B/op          1 allocs/op
BenchmarkLevenshtein/Chinese_Word-2                                      5000000               294 ns/op              48 B/op          1 allocs/op
BenchmarkLevenshtein/Chinese_Word-4                                      5000000               268 ns/op              48 B/op          1 allocs/op
BenchmarkLevenshtein/Chinese_Word-6                                      5000000               267 ns/op              48 B/op          1 allocs/op
BenchmarkLevenshtein/Chinese_Word-8                                      5000000               265 ns/op              48 B/op          1 allocs/op
BenchmarkLevenshtein/First_Word_include_Second_Word_-_2-2               10000000               174 ns/op              32 B/op          1 allocs/op
BenchmarkLevenshtein/First_Word_include_Second_Word_-_2-4               10000000               171 ns/op              32 B/op          1 allocs/op
BenchmarkLevenshtein/First_Word_include_Second_Word_-_2-6               10000000               172 ns/op              32 B/op          1 allocs/op
BenchmarkLevenshtein/First_Word_include_Second_Word_-_2-8               10000000               169 ns/op              32 B/op          1 allocs/op
BenchmarkLevenshtein/First_Word_Empty-2                                 50000000                24.7 ns/op             0 B/op          0 allocs/op
BenchmarkLevenshtein/First_Word_Empty-4                                 100000000               24.6 ns/op             0 B/op          0 allocs/op
BenchmarkLevenshtein/First_Word_Empty-6                                 50000000                24.4 ns/op             0 B/op          0 allocs/op
BenchmarkLevenshtein/First_Word_Empty-8                                 50000000                24.4 ns/op             0 B/op          0 allocs/op
BenchmarkLevenshtein/Second_Word_Empty-2                                50000000                25.3 ns/op             0 B/op          0 allocs/op
BenchmarkLevenshtein/Second_Word_Empty-4                                100000000               24.6 ns/op             0 B/op          0 allocs/op
BenchmarkLevenshtein/Second_Word_Empty-6                                50000000                27.4 ns/op             0 B/op          0 allocs/op
BenchmarkLevenshtein/Second_Word_Empty-8                                50000000                25.2 ns/op             0 B/op          0 allocs/op
BenchmarkLevenshtein/Equal-2                                            300000000                4.37 ns/op            0 B/op          0 allocs/op
BenchmarkLevenshtein/Equal-4                                            300000000                5.14 ns/op            0 B/op          0 allocs/op
BenchmarkLevenshtein/Equal-6                                            300000000                5.06 ns/op            0 B/op          0 allocs/op
BenchmarkLevenshtein/Equal-8                                            300000000                4.91 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/cloudfstrife/levenshtein     64.035s

```