package graph

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func FindLargestDisjointPaths(paths [][]string) [][]string {
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

func Getingdata() string {
	if len(os.Args) != 2 {
		fmt.Println("invalid syntax try to run\n\tgo run main.go <file name>")
		os.Exit(1)
	}
	fileName := os.Args[1]
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("the file does not exist")
		os.Exit(1)
	}

	return string(content)
}

func HandulFile(data []string) ([]string, []string, []string, []string) {
	var start, end []string
	var roomsNames, links []string
	foundstart, foundend := false, false
	for i := 0; i < len(data); i++ {
		if strings.HasPrefix(data[i], "#") {
			if data[i] == "##start" {
				if len(start) != 0 || foundstart {
					log.Fatal("invalid syntax")
				}
				foundstart = true
			} else if data[i] == "##end" {
				if len(end) != 0 || foundend {
					log.Fatal("invalid syntax")
				}
				foundend = true
			} else {
				continue
			}
		} else {
			if foundend && foundstart {
				log.Fatal("invalid syntax")
			}
			if foundstart {
				start = strings.Split(data[i], " ")
				roomsNames = append(roomsNames, strings.Split(data[i], " ")...)
				foundstart = false
			} else if foundend {
				end = strings.Split(data[i], " ")
				roomsNames = append(roomsNames, strings.Split(data[i], " ")...)
				foundend = false
			} else if strings.Contains(data[i], " ") {
				// fmt.Println(data[i])
				roomsNames = append(roomsNames, strings.Split(data[i], " ")...)
			} else if strings.Contains(data[i], "-") {
				links = append(links, strings.Split(data[i], "-")...)
			} else {
				fmt.Println(data[i])
				log.Fatal("invalid syntax")
			}
		}
	}
	return roomsNames, links, start, end
}
