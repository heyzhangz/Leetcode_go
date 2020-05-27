package question974

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
		input []int
		inputk int
		answer int
	}{
		{"test2", []int{2, -2, 2, -4}, 6, 2},
		{"test3", []int{-1, 2, 9}, 2, 2},
		{"test1", []int{4, 5, 0, -2, -3, 1}, 5, 7},
	}

	//test mine
	print("run test mine\n")
	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			if got := subarraysDivByK(test.input, test.inputk); test.answer != got {
				t.Errorf("subarraysDivByK() = %v, want %v", got, test.answer)
			}
		})
	}
}
