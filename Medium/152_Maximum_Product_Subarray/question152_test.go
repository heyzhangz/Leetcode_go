package question152

import "testing"

func TestQ(t *testing.T) {

	tests := []struct {
		testname string
		input []int
		answer int
	}{
		{"test5", []int{-2, -3}, 6},
		{"test5", []int{0, -2, 0}, 0},
		{"test4", []int{0}, 0},
		{"test1", []int{-2,0,-1}, 0},
		{"test3", []int{-2,0,2}, 2},
		{"test2", []int{2,3,-2,4}, 6},
	}

	//test stack
	print("run test stack\n")
	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			if got := maxProduct(test.input); test.answer != got {
				t.Errorf("maxProduct() = %v, want %v", got, test.answer)
			}
		})
	}
}
