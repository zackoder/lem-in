package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rooms struct {
	RoomList []*Room
}

type Room struct {
	Name       string
	Start, End bool
	Rooms      []*Room
}


func (rooms *Rooms) addRoomName(name string, start, end bool) {
	for _, r := range rooms.RoomList {
		if name == r.Name {
			return
		}
	}
	newRoom := &Room{Name: name, Start: start, End: end}
	rooms.RoomList = append(rooms.RoomList, newRoom)
}

func (rooms *Rooms) GetRoom(name string) *Room {
	for _, r := range rooms.RoomList {
		if r.Name == name {
			return r
		}
	}
	return nil
}

func (room *Rooms) addConnex(from, to string) {
	fromRoom := room.GetRoom(from)
	toRoom := room.GetRoom(to)
	if fromRoom != nil && toRoom != nil {
		for _, r := range fromRoom.Rooms {
			if r.Name == to {
				return
			}
		}
		fromRoom.Rooms = append(fromRoom.Rooms, toRoom)
		toRoom.Rooms = append(toRoom.Rooms, fromRoom)
	} else {
		fmt.Println("unknown room")
		os.Exit(0)
	}
}

func isDisjoint(path1, path2 []string) bool {
	rooms1 := make(map[string]bool)

	for _, room := range path1[1 : len(path1)-1] {
		rooms1[room] = true
	}
	for _, room := range path2[1 : len(path2)-1] {
		if rooms1[room] {
			return false
		}
	}
	return true
}

func findLargestDisjointPaths(paths [][]string) [][]string {
	var largestSet [][]string
	n := len(paths)

	var dfs func(idx int, current [][]string)
	dfs = func(idx int, current [][]string) {
		if idx == n {
			if len(current) > len(largestSet) {
				largestSet = append([][]string{}, current...)
			}
			return
		}

		canAdd := true
		for _, p := range current {
			if !isDisjoint(p, paths[idx]) {
				canAdd = false
				break
			}
		}
		if canAdd {
			dfs(idx+1, append(current, paths[idx]))
		}
		dfs(idx+1, current)
	}

	dfs(0, [][]string{})
	return largestSet
}

func (rooms *Rooms) Dfs(startRoom *Room, endRoom *Room) [][]string {

	var currentPath []string
	var allPaths [][]string
	visited := make(map[*Room]bool)
	var Dfshelper func(room *Room)
	Dfshelper = func(room *Room) {
		visited[room] = true
		currentPath = append(currentPath, room.Name)

		if room == endRoom {
			pathcopy := make([]string, len(currentPath))
			copy(pathcopy, currentPath)
			allPaths = append(allPaths, pathcopy)
		} else {
			for _, neighbors := range room.Rooms {
				if !visited[neighbors] {
					Dfshelper(neighbors)
				}
			}
		}
		currentPath = currentPath[:len(currentPath)-1]
		visited[room] = false
	}

	Dfshelper(startRoom)

	return allPaths
}

func main() {
	rooms := &Rooms{}

	fileData := getingdata()
	joinedData := strings.Join(fileData, "\n")
	data := strings.Split(joinedData, "\n")

	antsNUm, err := strconv.Atoi(data[0])
	if err != nil {
		fmt.Println(data[0], "is not a number")
		return
	}
	data = data[1:]

	var roomsNames, links []string
	var start, end []string
	for i := 0; i < len(data); i++ {
		if strings.Contains(data[i], "##start") {
			if i+1 <len(data) && strings.Contains(data[i+1], "#") {
				i++
			}
			start = strings.Split(data[i+1], " ")
		} else if strings.Contains(data[i], "##end") {
			if i+1 <len(data) && strings.Contains(data[i+1], "#") {
				i++
			}
			end = strings.Split(data[i+1], " ")
		} else if strings.Contains(data[i], " ") {
			roomsNames = append(roomsNames, strings.Split(data[i], " ")...)
		} else if strings.Contains(data[i], "-") {
			links = append(links, strings.Split(data[i], "-")...)
		} else if strings.Contains(data[i], "#") {
			continue
		} else {
			fmt.Println("invalid syntax")
			os.Exit(1)
		}
	}

	fmt.Println(antsNUm)

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

	for i := 0; i < len(roomsNames); i += 3 {
		if roomsNames[i] == string(start[0]) {
			rooms.addRoomName(roomsNames[i], true, false)
		} else if roomsNames[i] == string((end[0])) {
			rooms.addRoomName(roomsNames[i], false, true)
		} else {
			rooms.addRoomName(roomsNames[i], false, false)
		}
	}

	for len(links) >= 2 {
		rooms.addConnex(links[0], links[1])
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

		largestDisjointPaths = findLargestDisjointPaths(allPaths)
		fmt.Println(largestDisjointPaths)

	} else {
		fmt.Println("Start or end room not found!")
	}
}

func getingdata() []string {
	if len(os.Args) != 2 {
		fmt.Println("invalid syntax try to run\n\tgo run main.go <file name>")
		os.Exit(0)
	}
	fileName := os.Args[1]
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("the file does not exist")
	}
	
	return []string(strings.Split(string(content), "\n"))
}
