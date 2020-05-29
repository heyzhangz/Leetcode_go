package question198

import "testing"

func TestQ202(t *testing.T) {

	tests := []struct {
		testname string
		input []int
		answer int
	}{
		{"test4", []int{1, 1}, 1},
		{"test3", []int{}, 0},
		{"test1", []int{1, 2, 3, 1}, 4},
		{"test2", []int{2, 7, 9, 3, 1}, 12},
	}

	//test mine
	print("run test mine\n")
	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			if got := rob(test.input); test.answer != got {
				t.Errorf("rob() = %v, want %v", got, test.answer)
			}
		})
	}
}
