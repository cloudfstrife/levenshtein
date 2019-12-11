package levenshtein

import (
	"testing"
)

var testCases = map[string]struct {
	a    string
	b    string
	want int
}{
	"First Word Empty":                   {"", "levenshtein", 11},
	"Second Word Empty":                  {"levenshtein", "", 11},
	"UNEqual":                            {"levenshtein", "levenTESTCASE", 8},
	"Equal":                              {"hello", "hello", 0},
	"Chinese Word":                       {"你好世界", "晚安世界", 2},
	"First Word include Second Word - 1": {"你好世界", "世界", 2},
	"First Word include Second Word - 2": {"你好世界", "你好", 2},
	"Second Word include First Word - 1": {"世界", "晚安世界", 2},
	"Second Word include First Word - 2": {"晚安", "晚安世界", 2},
}

func TestLevenshtein(t *testing.T) {
	for name, testCase := range testCases {
		v := Levenshtein(testCase.a, testCase.b)
		if v != testCase.want {
			t.Errorf("%s Faild : want : %d got : %d", name, testCase.want, v)
		}
	}
}

func TestSimilarDegree(t *testing.T) {
	testCases := map[string]struct {
		a    string
		b    string
		want float64
	}{
		"First Word Empty":                   {"", "levenshtein", (1 - float64(11)/11)},
		"Second Word Empty":                  {"levenshtein", "", (1 - float64(11)/11)},
		"UNEqual":                            {"levenshtein", "levenTESTCASE", (1 - float64(8)/13)},
		"Equal":                              {"hello", "hello", (1 - float64(0)/5)},
		"Chinese Word":                       {"你好世界", "晚安世界", (1 - float64(2)/4)},
		"First Word include Second Word -1 ": {"你好世界", "世界", (1 - float64(2)/4)},
		"First Word include Second Word -2 ": {"你好世界", "你好", (1 - float64(2)/4)},
		"Second Word include First Word -1 ": {"世界", "晚安世界", (1 - float64(2)/4)},
		"Second Word include First Word -2 ": {"晚安", "晚安世界", (1 - float64(2)/4)},
	}

	for name, testCase := range testCases {
		v := SimilarDegree(testCase.a, testCase.b)

		if v != testCase.want {
			t.Errorf("%s Faild : want : %f got : %f", name, testCase.want, v)
		}
	}
}

func BenchmarkLevenshtein(b *testing.B) {
	for name, testCase := range testCases {
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Levenshtein(testCase.a, testCase.b)
			}
		})
	}
}
