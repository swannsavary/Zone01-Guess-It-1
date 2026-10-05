package tools

import (
	"math"
	"sort"
)

// renvoie la medianne de sorted
func Median() float64 {
	n := Len()
	if n%2 == 1 {
		return float64(sorted[n/2])
	}
	return float64(sorted[n/2-1]+sorted[n/2]) / 2
}

// renvoie la valeur absolue d'un entier
func Abs(nb int) int {
	if nb < 0 {
		return -nb
	}
	return nb
}

// médiane des |x - med|, sans retrier : on part de la médiane
// et on avance vers la gauche/droite en prenant le plus proche
func Mad(med float64) float64 {
	n := Len()
	l := sort.SearchInts(sorted, int(math.Ceil(med))) - 1
	r := l + 1
	var prev, cur float64
	for i := 0; i <= n/2; i++ {
		var d float64
		if l >= 0 && (r >= n || med-float64(sorted[l]) <= float64(sorted[r])-med) {
			d = med - float64(sorted[l])
			l--
		} else {
			d = float64(sorted[r]) - med
			r++
		}
		prev, cur = cur, d
	}
	if n%2 == 1 {
		return cur
	}
	return (prev + cur) / 2
}
