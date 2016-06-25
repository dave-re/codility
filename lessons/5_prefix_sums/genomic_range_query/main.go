package main

var (
	cMap = map[byte]int{
		0x41: 1,
		0x43: 2,
		0x47: 3,
		0x54: 4,
	}
	cList = []byte{0x41, 0x43, 0x47, 0x54}
)

func Solution(S string, P []int, Q []int) []int {
	sLen := len(S)
	pMap := map[byte][]int{
		0x41: make([]int, sLen),
		0x43: make([]int, sLen),
		0x47: make([]int, sLen),
		0x54: make([]int, sLen),
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
