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
			if len(p[k])+len(antQueues[k]) == min {
				antQueues[k] = append(antQueues[k], i)
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
	res := []string{}
	for _, v := range solution {
		for _, w := range v {
			res = append(res, w)
			fmt.Printf("%s ", w)
		}
		if len(v) == 0 {
			continue
		}
		fmt.Println()
	}
	fmt.Println(len(res))
}

func minLen(p [][]string, ants [][]int) int {
	min := math.MaxInt
	// fmt.Println(min)
	for i := range p {
		if min >= len(p[i])+len(ants[i]) {
			min = len(p[i]) + len(ants[i])
		}
		// fmt.Println(p[i], ants[i], "=", len(p[i])+len(ants[i]))
	}
	// fmt.Println(min)
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

func Lemin1(n int, p [][]string) ([][]string, int, int) {
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
	res := []string{}
	for _, v := range solution {
		res = append(res, v...)
	}
	return solution, len(res), len(solution)
}

func BestWay(chanl chan Data) [][]string {
	bestway := [][]string{}
	minclo := 0
	minrow := 0
	_ = minclo
	_ = minrow
	for r := range chanl {
		if len(bestway) == 0 {
			bestway = append(bestway, r.Realst...)
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
