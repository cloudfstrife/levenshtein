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
## benchmark


```
$ go test -v -bench="." -run="^$" -benchmem github.com/cloudfstrife/levenshtein
goos: windows
goarch: amd64
pkg: github.com/cloudfstrife/levenshtein
BenchmarkLevenshtein/UNEqual-8           1000000              1785 ns/op            1632 B/op         13 allocs/op
BenchmarkLevenshtein/Second_Word_include_First_Word_-_2-8                5000000               302 ns/op             224 B/op          4 allocs/op
BenchmarkLevenshtein/First_Word_Empty-8                                 100000000               23.6 ns/op             0 B/op          0 allocs/op
BenchmarkLevenshtein/Second_Word_Empty-8                                50000000                24.2 ns/op             0 B/op          0 allocs/op
BenchmarkLevenshtein/Equal-8                                            300000000                4.18 ns/op            0 B/op          0 allocs/op
BenchmarkLevenshtein/Chinese_Word-8                                      3000000               449 ns/op             368 B/op          6 allocs/op
BenchmarkLevenshtein/First_Word_include_Second_Word_-_1-8                5000000               350 ns/op             288 B/op          6 allocs/op
BenchmarkLevenshtein/First_Word_include_Second_Word_-_2-8                5000000               358 ns/op             288 B/op          6 allocs/op
BenchmarkLevenshtein/Second_Word_include_First_Word_-_1-8                5000000               305 ns/op             224 B/op          4 allocs/op
PASS
ok      github.com/cloudfstrife/levenshtein     16.946s

```