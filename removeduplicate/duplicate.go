package removeduplicate

import (
	"sort"
)

//RemoveDuplicate helps removinf duglicate number
func RemoveDuplicate(n []int) []int {
	setOfNumber := []int{}
	check := map[int]bool{}
	sort.Ints(n)

	for _, v := range n {
		if check[v] != true {
			check[v] = true
			setOfNumber = append(setOfNumber, v)
		}
	}
	return setOfNumber
}
