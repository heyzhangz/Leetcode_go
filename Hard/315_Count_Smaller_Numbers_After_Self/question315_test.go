package question315

import "testing"

func Test(t *testing.T) {

	tests := []struct {
		testname string
		input []int
	}{
		{"test1", []int{0, 1, 2}},
	}

	//test mine
	print("run test mine\n")
	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			print(countSmaller(test.input))
		})
	}
}
