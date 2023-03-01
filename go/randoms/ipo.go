package randoms

import (
	"container/heap"
	"sort"
)

type Project struct {
	profit  int
	capital int
}

type PriorityQueue []Project

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].profit > pq[j].profit
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(Project)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func FindMaximizedCapital(k int, w int, profits []int, capital []int) int {
	projects := make([]Project, len(profits))
	for i := 0; i < len(profits); i++ {
		projects[i] = Project{profit: profits[i], capital: capital[i]}
	}

	sort.Slice(projects, func(i, j int) bool {
		return projects[i].capital < projects[j].capital
	})

	pq := make(PriorityQueue, 0)
	i := 0

	for k > 0 {
		for i < len(projects) && projects[i].capital <= w {
			heap.Push(&pq, projects[i])
			i++
		}

		if len(pq) == 0 {
			break
		}

		p := heap.Pop(&pq).(Project)
		w += p.profit
		k--
	}

	return w
}
