package main

type stack []int

func (s *stack) pop() int {
	if len(*s) == 0 {
		return -1
	}
	data := s.peek()
	*s = (*s)[:len(*s)-1]
	return data
}

func (s *stack) peek() int {
	if len(*s) == 0 {
		return -1
	}
	data := (*s)[len(*s)-1]
	return data
}

func (s *stack) push(data int) {
	*s = append(*s, data)
}

func solution(order []int) int {
	var (
		boxNum      int = 1
		cnt, target int
		belt        *stack
	)

	belt = new(stack)

	for cnt < len(order) {
		target = order[cnt]

		if target > boxNum {
			belt.push(boxNum)
			boxNum++
			continue
		}

		switch target {
		case belt.peek():
			// from belt
			belt.pop()
		case boxNum:
			// from source
			boxNum++
		default:
			// no more move
			return cnt
		}
		cnt++
	}
	return cnt
}
