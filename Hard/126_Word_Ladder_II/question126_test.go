package question126

import "testing"

func TestQ(t *testing.T) {

	tests := []struct {
		testname string
		inputa string
		inputb string
		inputc []string
	}{
		{"test1", "hit", "cog", []string{"hot","dot","dog","lot","log","cog"}},
	}

	//test mine
	//print("run test mine\n")
	//for _, test := range tests {
	//	t.Run(test.testname, func(t *testing.T) {
	//		if got := largestRectangleArea(test.input); test.answer != got {
	//			t.Errorf("largestRectangleArea() = %v, want %v", got, test.answer)
	//		}
	//	})
	//}

	//test stack
	print("run test stack\n")
	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			got := findLadders(test.inputa, test.inputb, test.inputc)
			print(got)
		})
	}
}
