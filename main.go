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
	Counter    int
	X, Y       string
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
		fromRoom.Counter++
		toRoom.Counter++
		fromRoom.Rooms = append(fromRoom.Rooms, toRoom)
		toRoom.Rooms = append(toRoom.Rooms, fromRoom)
	}
}

// DFS Method to find all possible paths
// func (rooms *Rooms) dfs(startRoom *Room, endRoom *Room) [][]string {
// 	var allPaths [][]string
// 	var currentPath []string
// 	visited := make(map[*Room]bool)

// 	// Recursive DFS helper function
// 	var dfsHelper func(room *Room)
// 	dfsHelper = func(room *Room) {
// 		// Mark room as visited and add to current path
// 		visited[room] = true
// 		currentPath = append(currentPath, room.Name)

// 		// If we reached the end, save the path
// 		if room == endRoom {
// 			pathCopy := make([]string, len(currentPath))
// 			copy(pathCopy, currentPath)
// 			allPaths = append(allPaths, pathCopy)
// 		} else {
// 			// Recursively visit all adjacent rooms that haven't been visited
// 			for _, neighbor := range room.Rooms {
// 				if !visited[neighbor] {
// 					dfsHelper(neighbor)
// 				}
// 			}
// 		}

// 		// Backtrack: remove room from current path and mark as unvisited
// 		currentPath = currentPath[:len(currentPath)-1]
// 		visited[room] = false
// 	}

// 	// Start DFS from the start room
// 	dfsHelper(startRoom)

// 	return allPaths
// }

func (rooms *Rooms) Dfs(startRoom *Room, endRoom *Room) [][]string {
	var currentPath []string
	var allPaths [][]string
	visited := make(map[*Room]bool)
	var dfshelper func(room *Room)
	dfshelper = func(room *Room) {
		visited[room] = true
		currentPath = append(currentPath, room.Name)

		if room == endRoom {
			pathcopy := make([]string, len(currentPath))
			copy(pathcopy, currentPath)
			allPaths = append(allPaths, pathcopy)
		} else {
			for _, neigbors := range room.Rooms {
				if !visited[neigbors] {
					dfshelper(neigbors)
				}
			}
		}
		currentPath = currentPath[:len(currentPath)-1]
		visited[room] = false
	}

	dfshelper(startRoom)

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
			if strings.Contains(data[i+1], "#") {
				i++
			}
			start = strings.Split(data[i+1], " ")
		} else if strings.Contains(data[i], "##end") {
			if strings.Contains(data[i+1], "#") {
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

	var paths [][]string
	// Find and print all possible paths using DFS
	startRoom := rooms.GetRoom(string(start[0]))
	endRoom := rooms.GetRoom(string(end[0]))

	if startRoom != nil && endRoom != nil {
		paths = rooms.Dfs(startRoom, endRoom)
	} else {
		fmt.Println("Start or end room not found!")
	}
	fmt.Println(antsNUm)
	var ints []int
	counter := 0
	for _, path := range paths {
		for _, room := range path {
			counter += addCounters(rooms, room)
		}
		ints = append(ints, counter)
		counter = 0
	}
	for _, v := range paths {
		fmt.Println(v)
	}
	fmt.Println()
	com(paths, ints)
	// paths = append(paths[i:], paths...)
	// ints = append(ints[i:], ints...)
}

func com(paths [][]string, ints []int) {

	var validpaths = make(map[string][]string)


	

	// for i := 0; i < len(ints); i++ {
	// 	for j := i + 1; j < len(ints); j++ {
	// 		for _, path := range paths {
	// 			for index, room := range path {
	// 				if index == 0 || index == len(path)-1 {
	// 					continue
	// 				}
	// 				//fmt.Println(room)
	// 				str := strings.Join(paths[j], " ")
	// 				if !strings.Contains(str, room) {
	// 					if len(paths[i]) < len(paths[j]) || ints[i]-len(paths[i]) < ints[j]-len(paths[j]) {
	// 						str2 := strings.Join(paths[i], " ")
	// 						validpaths[str2] = paths[i]
	// 						break
	// 					}
	// 				} else if strings.Contains(str, room) {

	// 					str2 := strings.Join(paths[i], " ")
	// 					validpaths[str2] = paths[j]
	// 					validpaths[str] = paths[j]
	// 					break

	// 				}
	// 			}

	// 		}
	// 	}
	// }
	for _, v := range validpaths {
		fmt.Println(v)
	}
}

func addCounters(rooms *Rooms, name string) int {
	for _, r := range rooms.RoomList {
		if r.Name == name {
			return r.Counter
		}
	}
	return 0
}

/* func CheckingPaths(paths [][]string, ants int) {

	for _, v := range paths {
		fmt.Println(v)
	}

	var p = make(map[string][]string)
	var tempaths [][]string

	fmt.Println()
	fmt.Println(tempaths)
	for i := 0; i < len(paths)-1; i++ {
		tmp := strings.Join(paths[i], " ")
		tempaths = append(paths[:i], paths[i+1:]...)

	}

	fmt.Println()
	// for _,v := range p {
	// 	fmt.Println(v)
	// }

} */

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

	return []string{string(content)}
}
