package question63

import "testing"

func TestQ105(t *testing.T) {

	tests := []struct {
		testname string
		in [][]int
		answer int
	}{
		{"test1", [][]int{{0,0,0},{0,1,0},{0,0,0}}, 2},
	}

	//test mine
	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			if got := uniquePathsWithObstacles(test.in); test.answer != got {
				t.Errorf("uniquePathsWithObstacles() = %v, want %v", got, test.answer)
			}
		})
	}
}
