package question785

/*
	785. 判断二分图

	给定一个无向图graph，当这个图为二分图时返回true。
	如果我们能将一个图的节点集合分割成两个独立的子集A和B，并使图中的每一条边的两个节点一个来自A集合，一个来自B集合，我们就将这个图称为二分图。
	graph将会以邻接表方式给出，graph[i]表示图中与节点i相连的所有节点。每个节点都是一个在0到graph.length-1之间的整数。这图中没有自环和平行边： graph[i] 中不存在i，并且graph[i]中没有重复的值。

 */

/*
	思路: 没啥思路
	题解:
		染色法BFS\DFS还行
		并查集？
*/

const (
	UNTAGGED uint8 = 0
	RED uint8 = 1
	GREEN uint8 = 2
)

func isBipartite(graph [][]int) bool {

	colors := make([]uint8, len(graph))

	for i := range graph {
		if colors[i] == UNTAGGED {
			colors[i] = RED
			queue := []int{i}

			for len(queue) > 0 {
				cur := queue[0]
				queue = queue[1:]

				for _, j := range graph[cur] {
					if colors[j] == UNTAGGED {
						if colors[cur] == RED {
							colors[j] = GREEN
						} else {
							colors[j] = RED
						}
						queue = append(queue, j)
					} else {
						if colors[j] == colors[cur] {
							return false
						}
					}
				}
			}
		}
	}

	return true
}