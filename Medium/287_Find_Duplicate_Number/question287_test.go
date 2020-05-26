package question287

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
		answer int
	}{
		{"test1", []int{1, 3, 4, 2, 2}, 2},
		{"test2", []int{3, 1, 3, 4, 2}, 3},
	}

	//test mine
	print("run test mine\n")
	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			if got := findDuplicate(test.input); test.answer != got {
				t.Errorf("findDuplicate() = %v, want %v", got, test.answer)
			}
		})
	}
}
