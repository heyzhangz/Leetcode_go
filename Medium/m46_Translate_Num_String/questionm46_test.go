package questionm46

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
		input int
		answer int
	}{
		{"test2", 506, 1},
		{"test1", 12258, 5},
	}

	//test mine
	print("run test mine\n")
	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			if got := translateNum(test.input); test.answer != got {
				t.Errorf("translateNum() = %v, want %v", got, test.answer)
			}
		})
	}
}
