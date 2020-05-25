package question202

import "testing"

func TestQ202(t *testing.T) {

	tests := []struct {
		testname string
		input int
		answer bool
	}{
		{"test1", 19, true},
		{"test2", 2, false},
	}

	//test mine
	print("run test mine\n")
	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			if got := isHappy(test.input); test.answer != got {
				t.Errorf("isHappy() = %v, want %v", got, test.answer)
			}
		})
	}
}
