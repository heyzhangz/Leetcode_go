package question128

import "testing"

func TestQ(t *testing.T) {

	tests := []struct {
		testname string
		input []int
		answer int
	}{
		{"test2", []int{1, -8, 7, -2, -4, -4, 6, 3, -4, 0, -7, -1, 5, 1, -9, -3}, 6},
		{"test1", []int{100, 4, 200, 1, 3, 2}, 4},
	}

	//test mine
	//print("run test mine\n")
	//for _, test := range tests {
	//	t.Run(test.testname, func(t *testing.T) {
	//		if got := largestRectangleArea(test.input); test.answer != got {
	//			t.Errorf("largestRectangleArea() = %v, want %v", got, test.answer)
	//		}
	//	})
	//}

	//test stack
	print("run test stack\n")
	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			if got := longestConsecutive_UF(test.input); test.answer != got {
				t.Errorf("largestRectangleArea() = %v, want %v", got, test.answer)
			}
		})
	}
}
