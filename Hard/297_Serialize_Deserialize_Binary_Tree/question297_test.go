package question297

import "testing"

func Test(t *testing.T) {

	tests := []struct {
		testname string
		preorder []string
		inorder []string
		output string
	}{
		{"test2", []string{"1_0", "2_0", "2_1"}, []string{"2_0", "1_0", "2_1"}, "ok"},
	}

	//test mine
	print("run test mine\n")
	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			gotroot := buildTree(test.preorder, test.inorder)
			obj := Constructor()
			data := obj.serialize(gotroot)
			ans := obj.deserialize(data)
			print(ans)
		})
	}
}
