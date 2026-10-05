package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	tools "student/mathtools"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for {
		//si plus de donnees a lire alors break
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		//prend chaque ligne, la convertit en int et l'ajoute a sorted
		num, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			continue
		}
		tools.Insert(num)
		min, max := calculerPlage()

		fmt.Fprintf(writer, "%v %v\n", min, max)
		writer.Flush()
	}
}

func calculerPlage() (float64, float64) {
	m := tools.Median()
	if tools.Len() < 5 {
		return m * 0.8, m * 1.2
	}
	dif := 1.25
	d := tools.Mad(m) * 1.4826
	return m - dif*d, m + dif*d
}
