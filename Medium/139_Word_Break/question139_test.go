package question139

import "testing"

func TestQ5(t *testing.T) {

	tests := []struct {
		testname string
		input string
		b []string
		answer bool
	}{
		{"test2", "keetcode", []string{"keet", "code"}, true},
		{"test1", "aab", []string{"a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"}, false},
	}

	//test mine
	print("run test mine\n")
	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			if got := wordBreak(test.input, test.b); test.answer != got {
				t.Errorf("wordBreak() = %v, want %v", got, test.answer)
			}
		})
	}
}
