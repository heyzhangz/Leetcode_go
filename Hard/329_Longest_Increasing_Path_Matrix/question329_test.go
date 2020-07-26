package question329
import "testing"

func Test(t *testing.T) {

	tests := []struct {
		testname string
		input [][]int
	}{
		{"test1", [][]int{{9,9,4}, {6,6,8}, {2,1,1}}},
	}

	//test mine
	print("run test mine\n")
	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			print(longestIncreasingPath(test.input))
		})
	}
}
