package solution

const maxNum = 1000000000

type LeaderItem struct {
	IsLeader bool
	Value    int
}

func Solution(A []int) int {
	aLen := len(A)
	countMap := make(map[int]int, aLen)
	leaderArr := make([]LeaderItem, aLen)
	leaderCount := 0
	leaderVal := 0
	for i, v := range A {
		countMap[v]++
		if leaderCount < countMap[v] {
			leaderVal = v
			leaderCount = countMap[v]
		}
		if leaderCount > ((i + 1) / 2) {
			leaderArr[i] = LeaderItem{
				IsLeader: true,
				Value:    leaderVal,
			}
		}
	}

	count := 0
	for i, v := range A {
		countMap[v]--
		leader := leaderArr[i]
		if leader.IsLeader {
			if countMap[leader.Value] > ((aLen - (i + 1)) / 2) {
				count++
			}
		}
	}
	return count
}
