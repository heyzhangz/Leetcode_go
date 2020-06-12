package question15

import "testing"

func Test(t *testing.T) {

	tests := []struct {
		testname string
		input []int
		output [][]int
	}{
		{"test2", []int{-4,-2,-2,-2,0,1,2,2,2,3,3,4,4,6,6}, [][]int{{-1, 0, 1}, {0, 0, 0}, {-3, 1, 2}}},
		{"test1", []int{1, 0, -1, 0, 0, 2, 0, -3}, [][]int{{-1, 0, 1}, {0, 0, 0}, {-3, 1, 2}}},
	}

	//test mine
	print("run test mine\n")
	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			got := threeSum(test.input)
			print(got)
		})
	}
}
