package question1028

import "testing"

func Test(t *testing.T) {

	tests := []struct {
		testname string
		preorder string
		inorder []string
		output string
	}{
		{"test2", "1-2--3--4-5--6--7",[]string{"2_0", "1_0", "2_1"}, "ok"},
	}

	//test mine
	print("run test mine\n")
	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			gotroot := recoverFromPreorder(test.preorder)
			print(gotroot)
		})
	}
}
