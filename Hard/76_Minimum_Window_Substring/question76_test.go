package question76

import "testing"

func TestQ105(t *testing.T) {

	tests := []struct {
		testname string
		inputs string
		inputt string
		answer string
	}{
		{"test1", "ADOBECODEBANC", "ABC", "BANC"},
		{"test4", "bba", "ab", "ba"},
		{"test3", "a", "a", "a"},
		{"test2", "a", "aa", ""},
	}

	//test mine
	print("run test mine\n")
	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			if got := minWindow(test.inputs, test.inputt); test.answer != got {
				t.Errorf("minWindow() = %v, want %v", got, test.answer)
			}
		})
	}
}
