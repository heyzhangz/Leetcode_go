package question378

import "testing"

func Test(t *testing.T) {

	tests := []struct {
		testname string
		inputa [][]int
		inputb int
		answer int
	}{
		{"test1", [][]int{{1, 2}, {1, 3}}, 3, 2},
	}

	//test mine
	print("run test mine\n")
	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			if got := kthSmallest(test.inputa, test.inputb); test.answer != got {
				t.Errorf("kthSmallest() = %v, want %v", got, test.answer)
			}
		})
	}
}
