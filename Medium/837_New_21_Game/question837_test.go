package question837

import "testing"

func contain(arr []string, key string) bool {

	for _, a := range arr {
		if a == key {
			return true
		}
	}

	return false
}

func Test(t *testing.T) {

	tests := []struct {
		testname string
		inputN int
		inputK int
		inputW int
		answer float64
	}{
		{"test4", 10, 1, 10, 1.0},
	}

	//test mine
	print("run test mine\n")
	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			if got := new21Game(test.inputN, test.inputK, test.inputW); test.answer != got {
				t.Errorf("new21Game() = %v, want %v", got, test.answer)
			}
		})
	}
}
