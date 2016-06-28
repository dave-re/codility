package solution

import "sort"

type (
	Disc struct {
		L int
		R int
	}
	DiscArr []Disc
)

func (a DiscArr) Len() int           { return len(a) }
func (a DiscArr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a DiscArr) Less(i, j int) bool { return a[i].L < a[j].L }

func Solution(A []int) int {
	discArr := make(DiscArr, 0, len(A))
	for i := range A {
		discArr = append(discArr, Disc{
			L: i - A[i],
			R: i + A[i],
		})
	}

	sort.Sort(discArr)

	count := 0
	for i := range discArr {
		for _, d := range discArr[i+1:] {
			if discArr[i].L <= d.L && d.L <= discArr[i].R {
				count++
				if count > 10000000 {
					return -1
				}
			} else {
				break
			}
		}
	}
	return count
}
