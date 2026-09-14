package tools

import (
	"math"
	"sort"
)

// renvoie la moyenne d'un slice d'int
func Average(nbs []int) float64 {
	var result int
	for _, nb := range nbs {
		result += nb
	}
	return float64(result) / float64(len(nbs))
}

// renvoie la medianne d'un slice d'int
func Median(nbs []int) float64 {
	sorted := make([]int, len(nbs))
	copy(sorted, nbs)
	sort.Ints(sorted)
	if len(sorted)%2 == 0 {
		return float64(sorted[len(sorted)/2-1]+sorted[len(sorted)/2]) / 2
	}
	return float64(sorted[len(sorted)/2])
}

// renvoie la variance d'un slice d'int
func Variance(nbs []int) float64 {
	moyenne := Average(nbs)
	var somme float64
	for _, nb := range nbs {
		diff := float64(nb) - moyenne
		somme += diff * diff
	}
	return somme / float64(len(nbs))
}

// renvoi l'ecart type d'un slice d'int
func StandardDeviation(nbs []int) float64 {
	return math.Sqrt(Variance(nbs))
}

func MAD(nbs []int) float64 {
	m := Median(nbs)
	ecart := []int{}
	for _, nb := range nbs {
		ecart = append(ecart, Abs(nb-int(m)))
	}
	return Median(ecart)
}

func Abs(nb int) int {
	if nb < 0 {
		return -nb
	}
	return nb
}
