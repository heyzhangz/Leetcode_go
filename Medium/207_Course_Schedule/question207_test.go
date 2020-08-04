package question207

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
		inputa int
		inputb [][]int
		answer bool
	}{
		{"test1", 3, [][]int{{1, 0}, {2, 0}}, true},
	}

	//test mine
	print("run test mine\n")
	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			if got := canFinish(test.inputa, test.inputb); test.answer != got {
				t.Errorf("canFinish() = %v, want %v", got, test.answer)
			}
		})
	}
}
