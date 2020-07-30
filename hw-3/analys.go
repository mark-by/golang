package analys

import (
	"container/heap"
	"regexp"
	"strings"
)

type word struct {
	value    string
	priority int
}

type priorityQueue []*word

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	item := x.(*word)
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

//Analys возвращает 10 наиболее часто встречающихся слов в тексте
func Analys(text string) (result []string) {
	re := regexp.MustCompile(`[\p{L}\d_]+`)
	matches := re.FindAllString(text, -1)
	counter := map[string]int{}
	for _, match := range matches {
		counter[strings.ToLower(match)]++
	}
	pq := make(priorityQueue, 0, len(counter))
	for value, priority := range counter {
		pq = append(pq, &word{
			value:    value,
			priority: priority,
		})
	}
	heap.Init(&pq)
	max := 10
	if pq.Len() < max {
		max = pq.Len()
	}
	for i := 0; i < max; i++ {
		word := heap.Pop(&pq).(*word)
		result = append(result, word.value)
	}
	return result
}
