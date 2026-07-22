type Item struct {
	key int 
	value int
}
type MinHeap struct {
	data []Item
	capacity int
}

func (h *MinHeap) Peek() (Item, bool) {
	if len(h.data) == 0 {
		return Item{}, false
	}
	return h.data[0], true
}

func (h *MinHeap) Push(val Item) {
	if len(h.data) < h.capacity {
		h.data = append(h.data, val)
		h.siftUp(len(h.data) - 1)
		return
	}

	if val.value < h.data[0].value {
		return
	}

	h.data[0] = val
	h.siftDown(0)
}

func (h *MinHeap) siftUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if h.data[i].value > h.data[parent].value {
			break
		}

		h.data[i], h.data[parent] = h.data[parent], h.data[i]
		i = parent
	}
}

func (h *MinHeap) siftDown(i int) {
	n := len(h.data)
	for {
		left := 2 * i + 1
		right := 2 * i + 2
		smallest := i

		if left < n && h.data[left].value < h.data[smallest].value {
			smallest = left
		}

		if right < n && h.data[right].value < h.data[smallest].value {
			smallest = right
		}

		if smallest == i {
			break
		}

		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
		i = smallest
	}
}
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	if k <= 0 {
		return []int{}
	}

	resultHeap := MinHeap{data: []Item{}, capacity: k}
	for key, value := range m {
		node := Item{key: key, value: value}
		resultHeap.Push(node)
	}

	result := []int{}

	for _, data := range resultHeap.data {
		result = append(result, data.key)
	}

	return result
}
