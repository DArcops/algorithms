package array

import "sort"

type Int []int

func (a Int) Sort() {
	sort.Ints(a)
}

func (a Int) Search(x int) {
	sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
}
