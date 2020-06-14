package question1300

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
		input []int
		target int
		answer int
	}{
		{"test1", []int{4, 9, 3}, 10, 3},
		{"test2", []int{2, 3, 5}, 10, 5},
		{"test3", []int{60864, 25176, 27249, 21296, 20204}, 56803, 11361},
		{"test2", []int{1, 1, 1}, 50, 1},
	}

	//test mine
	print("run test mine\n")
	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			if got := findBestValue_dir(test.input, test.target); test.answer != got {
				t.Errorf("findBestValue() = %v, want %v", got, test.answer)
			}
		})
	}
}
