package main

import "testing"

func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		// Normal cases
		{"asaaaaaaa", 2},
		{"dfaaassdfdw", 3},

		// Edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbbbb", 1},
		{"abcabcabcd", 4},

		// Chinse support
		{"这是一串中文", 9},
		{"一二三二一", 5},
	}

	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubStr(tt.s)
		if actual != tt.ans {
			// 	t.Errorf("lengthOfNonRepeatingSubStr(%s);" +
			// "got %d, expected %d;", tt.s, actual,tt.ans)
			t.Errorf("got %d for input %s; "+
				"expected %d", actual, tt.s, tt.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	s := "撒打发打发手动阀手动阀手动阀撒旦啊阿斯顿"
	ans := 11

	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("got %d for input %s; " +
		"expected %d", actual, s, ans)
		}
	}
}
