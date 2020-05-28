package question394

import "testing"

func contain(arr []string, key string) bool {

	for _, a := range arr {
		if a == key {
			return true
		}
	}

	return false
}

func TestQ5(t *testing.T) {

	tests := []struct {
		testname string
		input string
		answer string
	}{
		{"test4", "3[a]2[b4[F]c]", "aaabFFFFcbFFFFc"},
		{"test3", "3[a2[c]]", "accaccacc"},
		{"test2", "3[a]2[bc]", "aaabcbc"},
		{"test1", "2[abc]3[cd]ef", "abcabccdcdcdef"},
	}

	//test mine
	print("run test mine\n")
	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			if got := decodeString(test.input); test.answer != got {
				t.Errorf("decodeString() = %v, want %v", got, test.answer)
			}
		})
	}
}
