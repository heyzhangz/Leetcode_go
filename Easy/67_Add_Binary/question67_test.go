package question67

import "testing"

func TestQ202(t *testing.T) {

	tests := []struct {
		testname string
		inputa string
		inputb string
		answer string
	}{
		{"test1", "1111", "1111", "11110"},
	}

	//test mine
	print("run test mine\n")
	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			if got := addBinary(test.inputa, test.inputb); test.answer != got {
				t.Errorf("addBinary() = %v, want %v", got, test.answer)
			}
		})
	}
}
