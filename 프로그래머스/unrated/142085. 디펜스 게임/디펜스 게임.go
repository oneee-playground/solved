package main

func solution(n int, k int, enemy []int) int {
	var (
		damage int
		used   int

		myHeap = new(maxheap)
	)

	for idx, enemyCnt := range enemy {
		myHeap.push(enemyCnt)

		damage += enemyCnt

		if damage > n {
			if used == k {
				return idx
			}
			used++

			damage -= myHeap.pop()
		}
	}
	return len(enemy)
}

type maxheap []int

func (m *maxheap) leaf(index int) bool { return index >= (len(*m)/2) && index <= len(*m) }

func (_ maxheap) parent(index int) int { return (index - 1) / 2 }

func (_ maxheap) leftchild(index int) int { return 2*index + 1 }

func (_ maxheap) rightchild(index int) int { return 2*index + 2 }

func (m *maxheap) swap(i, j int) { (*m)[i], (*m)[j] = (*m)[j], (*m)[i] }

func (m *maxheap) push(value int) {
	*m = append(*m, value)

	index := len(*m) - 1
	for (*m)[index] > (*m)[m.parent(index)] {
		m.swap(index, m.parent(index))
		index = m.parent(index)
	}
}

func (m *maxheap) pop() int {
	top := (*m)[0]

	(*m)[0] = (*m)[len(*m)-1]
	(*m) = (*m)[:len(*m)-1]

	m.downHeapify(0)

	return top
}

func (m *maxheap) downHeapify(current int) {
	if m.leaf(current) {
		return
	}

	largest := current
	leftIdx := m.leftchild(current)
	rightIdx := m.rightchild(current)

	if leftIdx < len(*m) && (*m)[leftIdx] > (*m)[largest] {
		largest = leftIdx
	}
	if rightIdx < len(*m) && (*m)[rightIdx] > (*m)[largest] {
		largest = rightIdx
	}
	if largest != current {
		m.swap(current, largest)
		m.downHeapify(largest)
	}
}
