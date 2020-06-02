package questionm64

import "testing"

func contain(arr []string, key string) bool {

	for _, a := range arr {
		if a == key {
			return true
		}
	}

	return false
}

func TestQ5(t *testing.T) {

	tests := []struct {
		testname string
		input int
		answer int
	}{
		{"test2", 1, 1},
		{"test3", 3, 6},
		{"test1", 9, 45},
	}

	//test mine
	print("run test mine\n")
	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			if got := sumNumsQ(test.input); test.answer != got {
				t.Errorf("sumNums() = %v, want %v", got, test.answer)
			}
		})
	}
}
