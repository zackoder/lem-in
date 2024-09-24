package graph

import (
	"fmt"
	"os"
)

type Rooms struct {
	RoomList []*Room
}

type Room struct {
	Name       string
	Start, End bool
	Rooms      []*Room
}

func (rooms *Rooms) AddRoomName(name string, start, end bool) {
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

func (room *Rooms) AddConnex(from, to string) {
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

func (rooms *Rooms) Dfs(startRoom *Room, endRoom *Room) [][]string {
	var allPaths [][]string
	visited := make(map[*Room]bool)
	DfsHelper(startRoom, endRoom, visited, []string{}, &allPaths)
	return allPaths
}

func DfsHelper(room *Room, endroom *Room, visited map[*Room]bool, currentPath []string, allPath *[][]string) {
	visited[room] = true
	currentPath = append(currentPath, room.Name)
	if room == endroom {
		pathcopy := make([]string, len(currentPath))
		copy(pathcopy, currentPath)
		*allPath = append(*allPath, pathcopy)
	} else {
		for _, neighbore := range room.Rooms {
			if !visited[neighbore] {
				// fmt.Println(visited[neighbore])
				DfsHelper(neighbore, endroom, visited, currentPath, allPath)
			}
		}
	}

	currentPath = currentPath[:len(currentPath)-1]
	visited[room] = false
}
