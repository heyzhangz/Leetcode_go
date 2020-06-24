package question16

import "testing"

func Test(t *testing.T) {

	tests := []struct {
		testname string
		input []int
		inputa int
		output int
	}{
		{"test2", []int{1, 1, 1, 0}, 100, 3},
	}

	//test mine
	print("run test mine\n")
	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			got := threeSumClosest(test.input, test.inputa)
			print(got)
		})
	}
}
