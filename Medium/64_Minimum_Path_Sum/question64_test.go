package question64

import "testing"

func Test(t *testing.T) {

	tests := []struct {
		testname string
		input [][]int
	}{
		{"test2", [][]int{{1, 2, 5}, {3, 2, 1}}},
	}

	//test mine
	print("run test mine\n")
	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			got := minPathSum(test.input)
			print(got)
		})
	}
}
