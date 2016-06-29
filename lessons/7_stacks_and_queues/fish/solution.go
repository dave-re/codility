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

func (s *Stack) Size() int {
	return s.top + 1
}

func Solution(A []int, B []int) int {
	upCount := 0
	downStack := NewStack(len(A))
	for i := range A {
		if B[i] == 1 {
			downStack.Push(A[i])
		} else {
			upFish := A[i]
			isAliveUpFish := true
			size := downStack.Size()
			for k := 0; k < size; k++ {
				downFish := downStack.Pop()
				if downFish > upFish {
					isAliveUpFish = false
					downStack.Push(downFish)
					break
				}
			}

			if isAliveUpFish {
				upCount++
			}
		}
	}

	return upCount + downStack.Size()
}
