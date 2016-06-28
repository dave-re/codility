package solution

// Solution is solution
func Solution(A []int) int {
	ret := A[0]
	arr := A[1:]
	for _, val := range arr {
		ret = ret ^ val
	}
	return ret
}

// func Solution(A []int) int {
// 	checkMap := make(map[int]bool, (len(A)/2 + 1))
// 	for _, val := range A {
// 		checkMap[val] = !checkMap[val]
// 	}
//
// 	for key, val := range checkMap {
// 		if val {
// 			return key
// 		}
// 	}
// 	return 0
// }
