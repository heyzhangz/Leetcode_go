package question4

import "testing"

func TestQ4(t *testing.T) {

	tests := []struct {
		testname string
		inputa []int
		inputb []int
		answer float64
	}{
		{"test6", []int{1}, []int{2, 3, 4, 5, 6, 7}, 4.0},
		{"test5", []int{}, []int{1}, 1.0},
		{"test3", []int{}, []int{1, 2, 3, 4, 5}, 3.0},
		{"test4", []int{}, []int{1, 2, 3, 4}, 2.5},
		{"test2", []int{1, 2}, []int{3, 4}, 2.5},
		{"test1", []int{1, 3}, []int{2}, 2.0},
	}

	//test mine
	print("run test mine\n")
	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			if got := findMedianSortedArrays(test.inputa, test.inputb); test.answer != got {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, test.answer)
			}
		})
	}
}
