// +build go1.8

package obj

import "sort"

func SortSlice(slice interface{}, less func(i, j int) bool) {
	sort.Slice(slice, less)
}
