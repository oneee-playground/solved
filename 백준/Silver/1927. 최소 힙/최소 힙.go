package main

import (
	"bufio"
	"fmt"
	"os"
)

// 최소 힙. 0번째 인덱스를 비워 놓는다고 가정
type MinHeap struct {
	elems []int // 동적 배열
}

// 삽입
func (h *MinHeap) Push(n int) {
	// 트리의 끝에 노드 삽입
	h.elems = append(h.elems, n)

	// 트리 조정
	h.upheapify(h.length() - 1)
}

// idx에 위치한 노드와 그 부모를 비교하여 swap하는 재귀적인 연산
func (h *MinHeap) upheapify(idx int) {
	// 부모 인덱스
	p := parent(idx)
	if p == 0 {
		// 현재 노드가 루트 노드이므로 끝
		return
	}

	if h.elems[p] <= h.elems[idx] {
		// 부모가 더 작거나 같으므로 끝
		return
	}

	// 교환 후 부모 인덱스에서 다시 검사
	h.swap(idx, p)
	h.upheapify(p)
}

// 삭제
func (h *MinHeap) Pop() int {
	if h.length() == 1 {
		// 엣지케이스
		return 0
	}

	root := h.elems[1]

	// 마지막 노드를 루트 자리에 넣기
	lastNode := h.elems[h.length()-1]
	h.elems[1] = lastNode

	// 배열 길이를 1 줄임
	h.elems = h.elems[:h.length()-1]

	// 트리 조정
	h.downheapify(1)

	return root
}

// idx에 위치한 노드와 그 자식들을 비교하여 swap하는 재귀적인 연산
func (h *MinHeap) downheapify(idx int) {
	// 왼쪽, 오른쪽 자식 인덱스
	lc, rc := leftChild(idx), rightChild(idx)

	if lc < h.length() && h.elems[lc] < h.elems[idx] {
		// 오른쪽 자식이 더 작기 때문에 스왑 후 이동
		if rc >= h.length() || h.elems[lc] < h.elems[rc] {
			// 왼쪽 자식이 오른쪽 자식보다 더 작기 때문에 스왑 후 이동
			h.swap(idx, lc)
			h.downheapify(lc)
			return
		}
	}
	if rc < h.length() && h.elems[rc] < h.elems[idx] {
		// 오른쪽 자식이 더 작기 때문에 스왑 후 이동
		h.swap(idx, rc)
		h.downheapify(rc)
		return
	}

	// 현재 노드가 자식보다 작음
}

func (h *MinHeap) swap(i, j int) {
	h.elems[i], h.elems[j] = h.elems[j], h.elems[i]
}

func (h *MinHeap) length() int { return len(h.elems) }

func parent(idx int) int     { return idx / 2 }
func leftChild(idx int) int  { return idx * 2 }
func rightChild(idx int) int { return idx*2 + 1 }

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	heap := MinHeap{[]int{0}} // 0번 인덱스 비우기

	var n, op int
	fmt.Fscan(r, &n)
	for ; n > 0; n-- {
		fmt.Fscan(r, &op)
		if op == 0 {
			fmt.Fprintln(w, heap.Pop())
			continue
		}
		heap.Push(op)
	}
}
