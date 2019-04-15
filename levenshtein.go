package levenshtein

import (
	"math"
)

//Levenshtein 两个字符串的编辑距离
func Levenshtein(a, b string) float64 {
	r1 := []rune(a)
	r2 := []rune(b)

	matrix := make([][]float64, 0, len(r1)+1)
	for i := 0; i < len(r1)+1; i++ {
		line := make([]float64, 0, len(r2)+1)
		for j := 0; j < len(r2)+1; j++ {
			if i == 0 {
				line = append(line, float64(j))
				continue
			}
			if j == 0 {
				line = append(line, float64(i))
				continue
			}
			var c float64
			if r1[i-1] != r2[j-1] {
				c = 1
			}
			val := math.Min(math.Min(matrix[i-1][j-1]+c, line[j-1]+1), matrix[i-1][j]+1)
			line = append(line, val)
		}
		matrix = append(matrix, line)
	}

	return matrix[len(r1)][len(r2)]
}

//SimilarDegree 字符串相似度
func SimilarDegree(a, b string) float64 {
	return 1 - Levenshtein(a, b)/math.Max(float64(len([]rune(a))), float64(len([]rune(b))))
}
