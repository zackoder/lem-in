package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	graph "graphs/Path"
)

func main() {
	rooms := graph.Rooms{}
	fileData := graph.Getingdata()
	data := strings.Split(fileData, "\n")
	antsNUm, err := strconv.Atoi(data[0])
	if err != nil {
		fmt.Println(data[0], "is not a number")
		return
	} else if antsNUm <= 0 {
		fmt.Println("put valid numbre")
		return
	}
	data = data[1:]
	var roomsNames, links []string
	var start, end []string
	roomsNames, links, start, end = graph.HandulFile(data)
	// fmt.Println(roomsNames, links, start, end)
	if len(start) == 0 {
		fmt.Println("you did not provide a start")
		os.Exit(0)
	}
	if len(end) == 0 {
		fmt.Println("you did not provide an end")
		os.Exit(0)
	}
	if len(roomsNames)%3 != 0 {
		fmt.Println("invalid syntax")
		os.Exit(0)
	}
	if start[0] == end[0] {
		return
	}
	for i := 0; i < len(roomsNames); i++ {
		if roomsNames[i] == start[0] {
			rooms.AddRoomName(roomsNames[i], true, false)
		} else if roomsNames[i] == end[0] {
			rooms.AddRoomName(roomsNames[i], false, true)
		} else {
			rooms.AddRoomName(roomsNames[i], false, false)
		}
	}
	for len(links) >= 2 {
		rooms.AddConnex(links[0], links[1])
		links = links[2:]
	}
	if len(links) > 0 {
		fmt.Println("invalid syntax")
		os.Exit(0)
	}

	startRoom := rooms.GetRoom(string(start[0]))
	endRoom := rooms.GetRoom(string(end[0]))
	// var largestDisjointPaths [][]string
	var allPaths [][]string
	//	var new [][]string
	if startRoom != nil && endRoom != nil {
		allPaths = rooms.Dfs(startRoom, endRoom)
		SortPath(allPaths)
		// if len(allPaths) > antsNUm {
		// 	new = graph.FindLargestDisjointPaths(allPaths[:antsNUm])
		// }
		// largestDisjointPaths = graph.FindLargestDisjointPaths(allPaths)
	} else {
		fmt.Println("Start or end room not found!")
		return
	}
	// largestDisjointPaths = DellSart(largestDisjointPaths)
	all := graph.AllPathDisjoin(allPaths)
	// for i, t := range all {
	// 	fmt.Println(i, t)
	// }

	var wg sync.WaitGroup
	chanl := make(chan graph.Data, len(all))
	TakeFunc := func(allPaths [][]string, n int) {
		defer wg.Done()
		res, col, row := graph.Lemin1(n, DellSart(allPaths))
		chanl <- graph.Data{
			Row:    row,
			Col:    col,
			Realst: res,
		}
	}

	wg.Add(len(all))
	for _, t := range graph.AllPathDisjoin(allPaths) {
		go TakeFunc(t, antsNUm)
	}
	wg.Wait()
	close(chanl)

	// for r := range chanl {
	// 	fmt.Println(r.Col)
	// 	fmt.Println(r.Row)
	// 	for t := range r.Realst {
	// 		fmt.Println(r.Realst[t])
	// 	}

	// 	fmt.Println("-----------------------------------------------------------")
	// }
	s := graph.BestWay(chanl)
	for t := range s {
		fmt.Println(strings.Join(s[t], " "))
	}

	// return
	// fmt.Println(largestDisjointPaths)
	// return
	// return
	// for _, t := range all {
	// 	//SortPath(t)
	// 	graph.Lemin(antsNUm, DellSart(t))
	// }
	// res := `L1-A0 L4-B0 L6-C0 L1-A1 L2-A0 L4-B1 L5-B0 L6-C1 L1-A2 L2-A1 L3-A0 L4-E2 L5-B1 L6-C2 L9-B0 L1-end L2-A2 L3-A1 L4-D2 L5-E2 L6-C3 L7-A0 L9-B1 L2-end L3-A2 L4-D3 L5-D2 L6-I4 L7-A1 L8-A0 L9-E2 L3-end L4-end L5-D3 L6-I5 L7-A2 L8-A1 L9-D2 L5-end L6-end L7-end L8-A2 L9-D3 L8-end L9-end`
	// fmt.Println(len(strings.Fields(res)))
}

func DellSart(s [][]string) [][]string {
	res := [][]string{}
	for _, t := range s {
		res = append(res, t[1:])
	}
	return res
}

func SortPath(p [][]string) {
	for i := 0; i < len(p); i++ {
		for j := i + 1; j < len(p); j++ {
			if len(p[i]) > len(p[j]) {
				p[i], p[j] = p[j], p[i]
			}
		}
	}
}
