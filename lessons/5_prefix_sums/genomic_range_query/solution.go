package solution

var (
	cMap = map[byte]int{
		'A': 1,
		'C': 2,
		'G': 3,
		'T': 4,
	}
	cList = []byte{'A', 'C', 'G', 'T'}
)

func Solution(S string, P []int, Q []int) []int {
	sLen := len(S)
	pMap := map[byte][]int{
		'A': make([]int, sLen),
		'C': make([]int, sLen),
		'G': make([]int, sLen),
		'T': make([]int, sLen),
	}

	// make a prefix sum
	for i := 0; i < len(S); i++ {
		sCh := S[i]
		for pCh := range pMap {
			prevNum := 0
			if i > 0 {
				prevNum = pMap[pCh][i-1]
			}

			if sCh == pCh {
				pMap[pCh][i] = prevNum + 1
			} else {
				pMap[pCh][i] = prevNum
			}
		}
	}

	pLen := len(P)
	rList := make([]int, pLen)
	for i := 0; i < pLen; i++ {
		if P[i] == Q[i] {
			rList[i] = cMap[S[P[i]]]
			continue
		}
		for _, ch := range cList {
			if P[i] == 0 && pMap[ch][0] > 0 {
				rList[i] = cMap[ch]
				break
			}

			pIdx := P[i] - 1
			if pIdx < 0 {
				pIdx = 0
			}
			if pMap[ch][Q[i]] > pMap[ch][pIdx] {
				rList[i] = cMap[ch]
				break
			}
		}
	}

	return rList
}
