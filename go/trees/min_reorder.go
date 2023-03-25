package trees

func MinReorder(n int, connections [][]int) int {
	// Build graph
	graph := make(map[int][]int)
	for _, c := range connections {
		graph[c[0]] = append(graph[c[0]], c[1])
		graph[c[1]] = append(graph[c[1]], -c[0])
	}
	
	// DFS
	var dfs func(int, int) int
	dfs = func(node, parent int) int {
		count := 0
		for _, child := range graph[node] {
			if child == parent {
				continue
			}
			if child < 0 {
				count += dfs(-child, node)
			} else {
				count += 1 + dfs(child, node)
			}
		}
		return count
	}
	
	return dfs(0, -1)
}
