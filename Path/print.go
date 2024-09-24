package graph

import (
	"fmt"
	"math"
)

func minLen(p [][]string, ants [][]int) int {
	min := math.MaxInt
	for i := range p {
		if min >= len(p[i])+len(ants[i]) {
			min = len(p[i]) + len(ants[i])
		}
	}
	return min
}

func MaxLen(p [][]string, ants [][]int) int {
	min := 0
	for i := range p {
		if min <= len(p[i])+len(ants[i]) {
			min = len(p[i]) + len(ants[i])
		}
	}
	return min
}

func Lemin(n int, p [][]string) ([][]string, int, int) {
	if len(p) == 0 {
		fmt.Println("there are no way from start to end")
		return nil, 0, 0
	}
	var antQueues [][]int = make([][]int, len(p))
	i := 1
	min := minLen(p, antQueues)
	// fmt.Println(n)
	for i <= n {
		for k := 0; k < len(p); k++ {
			if len(p[k])+len(antQueues[k]) == min {
				antQueues[k] = append(antQueues[k], i)
				break
			}
		}
		min = minLen(p, antQueues)
		i++
	}
	max := MaxLen(p, antQueues)
	var solution [][]string = make([][]string, max-1)
	for i := 0; i < len(p); i++ {
		for j, v := range antQueues[i] {
			for k, w := range p[i] {
				str := fmt.Sprintf("L%d-%s", v, w)
				solution[k+j] = append(solution[k+j], str)
			}
		}
	}
	res := 0
	for _, v := range solution {
		res += len(v)
	}
	return solution, res, len(solution)
}

func BestWay(chanl chan Data) [][]string {
	bestway := [][]string{}
	minclo := 0
	minrow := 0
	_ = minclo
	_ = minrow
	for r := range chanl {
		if len(bestway) == 0 {
			bestway = r.Realst
			minclo = r.Col
			minrow = r.Row
		} else {
			if r.Row < minrow {
				bestway = r.Realst
				minclo = r.Col
				minrow = r.Row
			} else if r.Row == minrow {
				if minclo > r.Col {
					bestway = r.Realst
					minclo = r.Col
					minrow = r.Row
				}
			}
		}
	}

	return bestway
}
