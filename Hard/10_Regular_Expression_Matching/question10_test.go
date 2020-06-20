package question10

import "testing"

func Test(t *testing.T) {

	tests := []struct {
		testname string
		ina string
		inb string
		output bool
	}{
		{"test11", "", "c*c*", true},
		{"test8", "a", "ab*", true},
		{"test10", "a", "aa", false},
		{"test9", "ab", ".*c", false},
		{"test4", "aab", "c*a*b", true},
		{"test1", "aaab", ".*", true},
		{"test3", "aa", "a*", true},
		{"test2", "aa", "a", false},
		{"test6", "aaa", "ab*a*c*a", true},
		{"test7", "aaa", "a*a", true},
		{"test5", "mississippi", "mis*is*p*.", false},
	}

	//test mine
	print("run test mine\n")
	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			if got := isMatch(test.ina, test.inb); test.output != got {
				t.Errorf("isMatch() = %v, want %v", got, test.output)
			}
		})
	}
}
