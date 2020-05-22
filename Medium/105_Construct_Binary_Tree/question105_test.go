package question105

import "testing"

func printPreTree(root *TreeNode) {

	if root != nil {
		print(root.Value)
		print( " ")
		printPreTree(root.Left)
		printPreTree(root.Right)
	}
}

func printInTree(root *TreeNode) {

	if root != nil {
		printInTree(root.Left)
		print(root.Value)
		print( " ")
		printInTree(root.Right)
	}
}

func TestQ105(t *testing.T) {

	tests := []struct {
		testname string
		preorder []int
		inorder []int
		output string
	}{
		{"test1", []int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}, "ok"},
		{"test2", []int{3}, []int{3}, "ok"},
		{"test3", []int{}, []int{}, "ok"},
	}

	//test mine
	print("run test mine\n")
	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			got := buildTree(test.preorder, test.inorder)
			print("preorder : ")
			printPreTree(got)
			println()
			print("inorder : ")
			printInTree(got)
			println()
		})
	}
}
