package graph

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isDisjoint(path1, path2 []string) bool {
	rooms1 := make(map[string]bool)
	if len(path2) == 2 && len(path1) == 2 {
		return false
	}
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
	startlink := false
	foundstart, foundend := false, false
	for i := 0; i < len(data); i++ {
		data[i] = strings.TrimRight(data[i], " ")
		if strings.HasPrefix(data[i], "#") {
			if data[i] == "##start" {
				if len(start) != 0 || foundstart {
					log.Fatal("ERROR: invalid data format")
				}
				foundstart = true
			} else if data[i] == "##end" {
				if len(end) != 0 || foundend {
					log.Fatal("ERROR: invalid data format")
				}
				foundend = true
			} else if strings.HasPrefix(data[i], "##") {
				log.Fatal(data[i], " is not commant")
				continue
			}
		} else {
			if foundend && foundstart {
				log.Fatal("ERROR: invalid data format")
			}
			if foundstart {
				start = strings.Split(data[i], " ")
				if len(start) != 3 {
					log.Fatal("ERROR: invalid data format")
				}
				LogEro(start[1])
				LogEro(start[2])
				roomsNames = append(roomsNames, strings.Split(data[i], " ")...)
				foundstart = false
			} else if foundend {
				end = strings.Split(data[i], " ")
				if len(end) != 3 {
					log.Fatal("ERROR: invalid data format")
				}
				LogEro(end[1])
				LogEro(end[2])
				roomsNames = append(roomsNames, strings.Split(data[i], " ")...)
				foundend = false
			} else if strings.Contains(data[i], " ") {
				if startlink {
					log.Fatal("ERROR: invalid data format")
				}
				split := strings.Split(data[i], " ")
				if len(split) != 3 {
					log.Fatal("ERROR: invalid data format")
				}
				LogEro(split[1])
				LogEro(split[2])
				roomsNames = append(roomsNames, split...)
			} else if strings.Contains(data[i], "-") {
				startlink = true
				split := strings.Split(data[i], "-")
				if len(split) != 2 {
					log.Fatal("ERROR: invalid data format")
				}
				if split[0] == "" || split[1] == "" {
					log.Fatal("ERROR: invalid data format")
				}
				links = append(links, strings.Split(data[i], "-")...)
			} else {
				log.Fatal("ERROR: invalid data format")
			}
		}
	}
	return roomsNames, links, start, end
}

func LogEro(s string) {
	_, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("you can't use string as coor")
	}
}
