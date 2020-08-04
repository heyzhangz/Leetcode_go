package question207
/*
	207. 课程表
	你这个学期必须选修 numCourse 门课程，记为 0 到 numCourse-1 。
	在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们：[0,1]
	给定课程总量以及它们的先决条件，请你判断是否可能完成所有课程的学习？
 */

/*
	思路: 拓扑排序吧
	题解:
*/

func canFinish(numCourses int, prerequisites [][]int) bool {

	inList := make([]int, numCourses)
	graph := make(map[int]map[int]bool, numCourses)

	for _, coupair := range prerequisites {
		inList[coupair[0]]++

		if _, ok := graph[coupair[1]]; !ok {
			graph[coupair[1]] = make(map[int]bool)
		}
		graph[coupair[1]][coupair[0]] = true
	}

	for true {

		delNode := 0
		findFlag := false
		for i, inD := range inList {
			if inD == 0 {
				delNode = i
				inList[i] = -1
				findFlag = true
				break
			}
		}

		if !findFlag {
			break
		}

		for n := range graph[delNode] {
			inList[n]--
		}
	}

	for _, n := range inList {
		if n != -1 {
			return false
		}
	}

	return true
}