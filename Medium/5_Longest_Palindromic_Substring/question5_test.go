package question5

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
		inputs string
		answer []string
	}{
		{"test1", "", []string{""}},
		{"test2", "ababab", []string{"ababa", "babab"}},
		{"test3", "aaaaaa", []string{"aaaaaa"}},
	}

	//test mine
	print("run test mine\n")
	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			if got := longestPalindrome(test.inputs); !contain(test.answer, got) {
				t.Errorf("longestPalindrome() = %v, want %v", got, test.answer)
			}
		})
	}

	//test DP
	print("run test DP\n")
	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			if got := longestpalindromeDp(test.inputs); !contain(test.answer, got) {
				t.Errorf("longestPalindrome_DP() = %v, want %v", got, test.answer)
			}
		})
	}
}
