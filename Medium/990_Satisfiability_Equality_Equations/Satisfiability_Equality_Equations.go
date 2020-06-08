package question990

/*
	990. 等式方程的可满足性
	给定一个由表示变量之间关系的字符串方程组成的数组，每个字符串方程 equations[i] 的长度为 4，并采用两种不同的形式之一："a==b" 或 "a!=b"。在这里，a 和 b 是小写字母（不一定不同），表示单字母变量名。
	只有当可以将整数分配给变量名，以便满足所有给定的方程时才返回 true，否则返回 false。 
 */

/*
	思路: 第一思路是并查集 只要先判断==构造并查集，然后判断!=的父节点是不是一个就好
	题解: 思路一样
*/

func isEquation(a byte, b byte, union map[byte]byte) bool {

	if findRoot(union, a) == findRoot(union, b) {
		return true
	}

	return false
}

func findRoot(union map[byte]byte, leaf byte) byte {

	ans := leaf
	for ; ans != union[ans]; {
		ans = union[ans]
	}

	return ans
}

func equationsPossible(equations []string) bool {

	var uneqlist []int
	unionfind := make(map[byte]byte, len(equations))

	for i, n := range equations {

		a := n[0]
		b := n[3]
		if _, ok := unionfind[a]; !ok {
			unionfind[a] = a
		}
		if _, ok := unionfind[b]; !ok {
			unionfind[b] = b
		}

		if n[1] == '!' {
			uneqlist = append(uneqlist, i)
		} else {
			roota := findRoot(unionfind, a)
			rootb := findRoot(unionfind, b)

			unionfind[rootb] = roota
		}
	}

	for _, i := range uneqlist {
		a := equations[i][0]
		b := equations[i][3]

		if isEquation(a, b, unionfind) {
			return false
		}
	}

	return true
}

//题解的并查集 利用了26个字母 不好扩展感觉
//func equationsPossible(equations []string) bool {
//	parent := make([]int, 26)
//	for i := 0; i < 26; i++ {
//		parent[i] = i
//	}
//	for _, str := range equations {
//		if str[1] == '=' {
//			index1 := int(str[0] - 'a')
//			index2 := int(str[3] - 'a')
//			union(parent, index1, index2)
//		}
//	}
//
//	for _, str := range equations {
//		if str[1] == '!' {
//			index1 := int(str[0] - 'a')
//			index2 := int(str[3] - 'a')
//			if find(parent, index1) == find(parent, index2) {
//				return false
//			}
//		}
//	}
//	return true
//}
//
//func union(parent []int, index1, index2 int) {
//	parent[find(parent, index1)] = find(parent, index2)
//}
//
//func find(parent []int, index int) int {
//	for parent[index] != index {
//		parent[index] = parent[parent[index]]
//		index = parent[index]
//	}
//	return index
//}