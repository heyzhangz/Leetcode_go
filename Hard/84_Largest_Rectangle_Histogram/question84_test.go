package question84

import "testing"

func TestQ105(t *testing.T) {

	tests := []struct {
		testname string
		input []int
		answer int
	}{
		{"test2", []int{5, 4, 1, 2}, 8},
		{"test1", []int{2, 1, 5, 6, 2, 3}, 10},
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
			if got := largestrectangleareaS(test.input); test.answer != got {
				t.Errorf("largestRectangleArea() = %v, want %v", got, test.answer)
			}
		})
	}
}
