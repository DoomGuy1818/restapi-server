package main

import (
	"sort"

	
)

func main() {
TwoOldestAges([]int{1, 7, 8, 4})
}

func TwoOldestAges(ages []int) [2]int {
	sort.Ints(ages)
	oldestAge := ages[len(ages)-2 : len(ages)-1]
	
	return [2]int(oldestAge)
}