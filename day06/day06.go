package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var grid map[string]bool
var newGrid map[string]int

func getPoint(val string) (int, int) {
	vals := strings.Split(val, ",")
	x, err := strconv.Atoi(vals[0])
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.Atoi(vals[1])
	if err != nil {
		log.Fatal(err)
	}
	return x, y
}

func modifyGrid(start, end string, f func(string) bool) {
	sx, sy := getPoint(start)
	ex, ey := getPoint(end)
	for i := sx; i <= ex; i++ {
		for j := sy; j <= ey; j++ {
			val := fmt.Sprintf("%d,%d", i, j)
			grid[val] = f(val)
		}
	}
}

func parse(line string) {
	words := strings.Split(line, " ")
	if line[:7] == "turn on" {
		modifyGrid(words[2], words[4], func(val string) bool {
			return true
		})
	} else if line[:8] == "turn off" {
		modifyGrid(words[2], words[4], func(val string) bool {
			return false
		})
	} else if line[:6] == "toggle" {
		modifyGrid(words[1], words[3], func(val string) bool {
			cur, ok := grid[val]
			if !ok {
				return true
			}
			return !cur
		})
	}
}

func modifyGridNew(start, end string, f func(string) int) {
	sx, sy := getPoint(start)
	ex, ey := getPoint(end)
	for i := sx; i <= ex; i++ {
		for j := sy; j <= ey; j++ {
			val := fmt.Sprintf("%d,%d", i, j)
			newGrid[val] = f(val)
		}
	}
}

func parseNew(line string) {
	words := strings.Split(line, " ")
	if line[:7] == "turn on" {
		modifyGridNew(words[2], words[4], func(val string) int {
			cur, ok := newGrid[val]
			if !ok {
				return 1
			}
			return cur + 1
		})
	} else if line[:8] == "turn off" {
		modifyGridNew(words[2], words[4], func(val string) int {
			cur, ok := newGrid[val]
			if !ok {
				return 0
			}
			if cur > 0 {
				return cur - 1
			}
			return 0
		})
	} else if line[:6] == "toggle" {
		modifyGridNew(words[1], words[3], func(val string) int {
			cur, ok := newGrid[val]
			if !ok {
				return 2
			}
			return cur + 2
		})
	}
}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	grid = make(map[string]bool)
	newGrid = make(map[string]int)
	var lines []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for _, line := range lines {
		parse(line)
	}
	count := 0
	for _, val := range grid {
		if val {
			count++
		}
	}
	fmt.Println(count)
	for _, line := range lines {
		parseNew(line)
	}
	count = 0
	for _, val := range newGrid {
		count += val
	}
	fmt.Println(count)
}
