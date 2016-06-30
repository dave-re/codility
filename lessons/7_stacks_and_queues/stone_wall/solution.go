package solution

type Stack struct {
	items []int
	top   int
}

func NewStack(maxSize int) *Stack {
	return &Stack{
		items: make([]int, maxSize),
		top:   -1,
	}
}

func (s *Stack) Push(v int) {
	s.top++
	s.items[s.top] = v
}

func (s *Stack) Pop() int {
	v := s.items[s.top]
	s.top--
	return v
}

func (s *Stack) Peek() int {
	return s.items[s.top]
}

func (s *Stack) Size() int {
	return s.top + 1
}

func Solution(H []int) int {
	count := 0
	st := NewStack(len(H))

	st.Push(H[0])
	for _, val := range H[1:] {
		stSize := st.Size()
		for i := 0; i < stSize; i++ {
			sVal := st.Peek()
			if sVal > val {
				count++
				st.Pop()
			} else if sVal == val {
				st.Pop()
			} else {
				break
			}
		}
		st.Push(val)
	}

	return count + st.Size()
}
