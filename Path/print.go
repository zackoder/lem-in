package graph

import (
	"fmt"
	"math"
)

func Lemin(n int, p [][]string) {
	if len(p) == 0 {
		fmt.Println("there are no way from start to end")
		return
	}
	var antQueues [][]int = make([][]int, len(p))
	i := 1
	min := minLen(p, antQueues)
	// fmt.Println(n)
	for i <= n {
		for k := 0; k < len(p); k++ {
			if len(p[k])+len(antQueues[k]) <= min {
				// fmt.Println(min, "----\n", len(antQueues[k]), len(p[k]))
				// fmt.Println(antQueues)
				antQueues[k] = append(antQueues[k], i)
				// fmt.Println(antQueues[k])
				break
			}
		}
		min = minLen(p, antQueues)

		i++
	}
	// return
	max := MaxLen(p, antQueues)
	// fmt.Println(max)
	var solution [][]string = make([][]string, max-1)
	for i := 0; i < len(p); i++ {
		for j, v := range antQueues[i] {
			for k, w := range p[i] {
				str := fmt.Sprintf("L%d-%s", v, w)
				solution[k+j] = append(solution[k+j], str)
			}
		}
	}
	// fmt.Println(solution)
	for _, v := range solution {
		for _, w := range v {
			fmt.Printf("%s ", w)
		}
		if len(v) == 0 {
			continue
		}
		fmt.Println()
	}
}

func minLen(p [][]string, ants [][]int) int {
	min := math.MaxInt
	// fmt.Println(min)
	for i := range p {
		if min >= len(p[i])+len(ants[i]) {
			min = len(p[i]) + len(ants[i])
		}
	}
	// fmt.Println(min)
	return min
}

func MaxLen(p [][]string, ants [][]int) int {
	min := 0
	for i := range p {
		if min < len(p[i])+len(ants[i]) {
			min = len(p[i]) + len(ants[i])
		}
	}
	return min
}
