package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	paper := 0
	ribbon := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		sides := strings.Split(line, "x")
		l, err := strconv.Atoi(sides[0])
		if err != nil {
			log.Fatal(err)
		}
		w, err := strconv.Atoi(sides[1])
		if err != nil {
			log.Fatal(err)
		}
		h, err := strconv.Atoi(sides[2])
		if err != nil {
			log.Fatal(err)
		}
		lwh := []int{l, w, h}
		sort.Ints(lwh)
		l = lwh[0]
		w = lwh[1]
		h = lwh[2]
		paper += l*w*3 + w*h*2 + l*h*2
		ribbon += l*2 + w*2 + l*w*h
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(paper)
	fmt.Println(ribbon)
}
