package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	graph "graphs/Path"
)

func main() {
	rooms := graph.Rooms{}
	fileData := graph.Getingdata()
	fileData = strings.TrimRight(fileData, "\n")
	data := strings.Split(fileData, "\n")
	antsNUm, err := strconv.Atoi(data[0])
	if err != nil {
		fmt.Println("invalid number of Ants")
		return
	} else if antsNUm <= 0 || antsNUm > 32000 {
		fmt.Println("you have right in range 1 and 32000")
		return
	}
	data = data[1:]
	var roomsNames, links []string
	var start, end []string
	roomsNames, links, start, end = graph.HandulFile(data)
	if len(start) == 0 {
		fmt.Println("no start room found")
		os.Exit(0)
	}
	if len(end) == 0 {
		fmt.Println("no end room found")
		os.Exit(0)
	}
	if len(roomsNames)%3 != 0 {
		fmt.Println("ERROR: invalid data format")
		os.Exit(0)
	}
	if start[0] == end[0] {
		fmt.Println("you are in end")
		return
	}
	for i := 0; i < len(roomsNames); i += 3 {
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
		fmt.Println("ERROR: invalid data format")
		os.Exit(0)
	}

	startRoom := rooms.GetRoom(string(start[0]))
	endRoom := rooms.GetRoom(string(end[0]))
	var allPaths [][]string
	if startRoom != nil && endRoom != nil {
		allPaths = rooms.Dfs(startRoom, endRoom)
		SortPath(allPaths)
	} else {
		fmt.Println("Start or end room not found!")
		return
	}
	if len(allPaths) == 0 {
		log.Fatalf("there are no paths between start and end")
	}
	all := graph.AllPathDisjoin(allPaths)
	var wg sync.WaitGroup
	chanl := make(chan graph.Data, len(all))
	TakeFunc := func(allPaths [][]string, n, index int) {
		defer wg.Done()
		res, col, row := graph.Lemin(n, DellSart(allPaths))
		chanl <- graph.Data{
			Row:    row,
			Col:    col,
			Realst: res,
			Index:  index,
		}
	}
	wg.Add(len(all))
	for i, t := range all {
		go TakeFunc(t, antsNUm, i)
	}
	wg.Wait()
	close(chanl)
	s, _ := graph.BestWay(chanl)
	fmt.Println(fileData)
	fmt.Println()
	for t := range s {
		if len(s[t]) == 0 {
			continue
		}
		fmt.Println(strings.Join(s[t], " "))
	}
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
