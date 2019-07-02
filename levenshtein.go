package levenshtein

import (
	"math"
)

//Levenshtein 两个字符串的编辑距离
func Levenshtein(a, b string) int {
	if a == b {
		return 0
	}

	r1 := []rune(a)
	r2 := []rune(b)

	if len(r1) == 0 {
		return len(r2)
	}

	if len(r2) == 0 {
		return len(r1)
	}

	line := make([]int, 0, len(r2)+1)
	nook := 0
	for i := 0; i < len(r1)+1; i++ {
		for j := 0; j < len(r2)+1; j++ {
			if i == 0 {
				line = append(line, j)
				continue
			}
			if j == 0 {
				nook = i - 1
				line[0] = i
				continue
			}
			oval := line[j]

			var v int
			if r1[i-1] != r2[j-1] {
				v = 1
			}
			line[j] = min(nook+v, line[j-1]+1, oval+1)
			nook = oval
		}
	}

	return line[len(r2)]
}

//SimilarDegree 字符串相似度
func SimilarDegree(a, b string) float64 {
	return 1 - float64(Levenshtein(a, b))/math.Max(float64(len([]rune(a))), float64(len([]rune(b))))
}

//min 获取小值
func min(a ...int) int {
	r := a[0]
	for _, v := range a {
		if r > v {
			r = v
		}
	}
	return r
}
