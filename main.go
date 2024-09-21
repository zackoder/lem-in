package main

import (
	"fmt"
	graph "graphs/Path"
	"os"
	"strconv"
	"strings"
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
	//fmt.Println(roomsNames, links, start, end)
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
	var largestDisjointPaths [][]string
	if startRoom != nil && endRoom != nil {
		allPaths := rooms.Dfs(startRoom, endRoom)
		largestDisjointPaths = graph.FindLargestDisjointPaths(allPaths)

	} else {
		fmt.Println("Start or end room not found!")
		return
	}
	largestDisjointPaths = DellSart(largestDisjointPaths)
	if len(largestDisjointPaths) == 0 {
		fmt.Println("there are no way from start to end")
		return
	}
	graph.Lemin(antsNUm, largestDisjointPaths)
}

func DellSart(s [][]string) [][]string {
	res := [][]string{}
	for _, t := range s {
		res = append(res, t[1:])
	}
	return res
}
