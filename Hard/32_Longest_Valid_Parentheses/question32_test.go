package question32

import "testing"

func Test(t *testing.T) {

	tests := []struct {
		testname string
		in string
		output int
	}{
		{"test1", "(()", 2},
		{"test2", ")()())", 4},
	}

	//test mine
	print("run test mine\n")
	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			if got := longestValidParentheses(test.in); test.output != got {
				t.Errorf("longestValidParentheses() = %v, want %v", got, test.output)
			}
		})
	}
}
