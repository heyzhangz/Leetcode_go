package question739

import "testing"

func equal(a []int, b []int) bool {

	for i, n := range a {
		if b[i] != n {
			return false
		}
	}

	return true
}

func Test(t *testing.T) {

	tests := []struct {
		testname string
		input []int
		answer []int
	}{
		{"test1", []int{73,74,75,71,69,72,76,73}, []int{1,1,4,2,1,1,0,0}},
	}

	//test mine
	print("run test mine\n")
	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			if got := dailyTemperatures(test.input); !equal(test.answer, got) {
				t.Errorf("dailyTemperatures() = %v, want %v", got, test.answer)
			}
		})
	}
}
