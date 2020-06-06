package questionm29

import "testing"

func Test(t *testing.T) {

	tests := []struct {
		testname string
		input [][]int
		answer []int
	}{
		{"test1", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
	}

	//test mine
	print("run test mine\n")
	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			if got := spiralOrder(test.input); isArrSame(test.answer, got) {
				t.Errorf("isHappy() = %v, want %v", got, test.answer)
			}
		})
	}
}

func isArrSame(A []int, B []int) bool {

	lenA := len(A)
	lenB := len(B)

	if lenA != lenB {
		return false
	}

	for i, n := range A {

		if n != B[i] {
			return false
		}
	}

	return true
}
