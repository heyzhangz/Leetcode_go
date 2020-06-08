package question990

import "testing"

func contain(arr []string, key string) bool {

	for _, a := range arr {
		if a == key {
			return true
		}
	}

	return false
}

func Test(t *testing.T) {

	tests := []struct {
		testname string
		input []string
		answer bool
	}{
		{"test2", []string{"c==c", "b==d", "z!=x"}, true},
		{"test1", []string{"a==b", "b!=c", "c==a"}, false},
	}

	//test mine
	print("run test mine\n")
	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			if got := equationsPossible(test.input); test.answer != got {
				t.Errorf("equationsPossible() = %v, want %v", got, test.answer)
			}
		})
	}
}
