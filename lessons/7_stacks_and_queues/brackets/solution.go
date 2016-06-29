package solution

type Stack struct {
	items []byte
	top   int
}

func NewStack(maxSize int) *Stack {
	return &Stack{
		items: make([]byte, maxSize),
		top:   -1,
	}
}

func (s *Stack) Push(v byte) {
	s.top++
	s.items[s.top] = v
}

func (s *Stack) Pop() byte {
	v := s.items[s.top]
	s.top--
	return v
}

func (s *Stack) Size() int {
	return s.top + 1
}

var (
	lbSet = map[byte]bool{
		'(': true,
		'{': true,
		'[': true,
	}
	rbSet = map[byte]bool{
		')': true,
		'}': true,
		']': true,
	}
	pairSet = map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}
)

func Solution(S string) int {
	if len(S) == 0 {
		return 1
	}
	stack := NewStack(len(S))
	for i := range S {
		c := S[i]
		if lbSet[c] {
			stack.Push(c)
		} else {
			if (stack.Size() == 0) || (pairSet[stack.Pop()] != c) {
				return 0
			}
		}
	}

	if stack.Size() == 0 {
		return 1
	}
	return 0
}
