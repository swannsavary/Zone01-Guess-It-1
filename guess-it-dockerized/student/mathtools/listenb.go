package tools

import "sort"

var sorted []int // toujours triée

// insère un entier dans sorted en gardant l'ordre
func Insert(v int) {
	i := sort.SearchInts(sorted, v)
	sorted = append(sorted, 0)
	copy(sorted[i+1:], sorted[i:])
	sorted[i] = v
}

// renvoie la longueur de sorted
func Len() int {
	return len(sorted)
}
