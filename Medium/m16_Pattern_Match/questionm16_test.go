package questionm16

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
		inputa string
		inputb string
		answer bool
	}{
		{"test1", "abba", "dogcatcatdog", true},
		{"test5", "bb", "thuhrh", false},
		{"test3", "aaaa", "dogcatcatdog", false},
		{"test2", "abba", "dogcatcatfish", false},
		{"test4", "abba", "dogdogdogdog", true},
	}

	//test mine
	print("run test mine\n")
	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			if got := patternMatching(test.inputa, test.inputb); test.answer != got {
				t.Errorf("patternMatching() = %v, want %v", got, test.answer)
			}
		})
	}
}
