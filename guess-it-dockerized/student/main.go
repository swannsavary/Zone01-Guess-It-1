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
	var allnb []int
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break // EOF, plus de données
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			continue
		}
		allnb = append(allnb, num)
		min, max := calculerPlage(allnb)

		fmt.Fprintf(writer, "%v %v\n", min, max)
		writer.Flush() // important : flush à chaque tour, pas juste à la fin
	}
}

func calculerPlage(allnb []int) (float64, float64) {
	if len(allnb) < 5 {
		moyenne := tools.Median(allnb)
		return moyenne * 0.8, moyenne * 1.2
	}
	dif := 0.40
	plagenb := 30
	fenetre := allnb
	if len(allnb) > plagenb {
		fenetre = allnb[len(allnb)-plagenb:]
	}
	moy := tools.Median(fenetre)
	mad := tools.MAD(fenetre) * 1.4826
	return moy - dif*mad, moy + dif*mad
}
