package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func coord(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	ox := 0
	oy := 0
	visited := map[string]struct{}{
		coord(0, 0): struct{}{},
	}
	for scanner.Scan() {
		line := scanner.Text()
		for _, c := range line {
			switch string(c) {
			case "^":
				oy++
			case "v":
				oy--
			case ">":
				ox++
			case "<":
				ox--
			}
			visited[coord(ox, oy)] = struct{}{}
		}
	}
	fmt.Println(len(visited))

	f.Seek(0, 0)
	scanner = bufio.NewScanner(f)
	realX := 0
	realY := 0
	roboX := 0
	roboY := 0
	visited = map[string]struct{}{
		coord(0, 0): struct{}{},
	}
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		for _, c := range line {
			var x, y *int
			if count%2 == 0 {
				x, y = &realX, &realY
			} else {
				x, y = &roboX, &roboY
			}

			switch string(c) {
			case "^":
				*y++
			case "v":
				*y--
			case ">":
				*x++
			case "<":
				*x--
			}
			visited[coord(*x, *y)] = struct{}{}
			count++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(visited))
}
