package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	floor := 0
	firstBasement := -1
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		for index, char := range line {
			if string(char) == "(" {
				floor++
			} else if string(char) == ")" {
				floor--
				if floor == -1 {
					if firstBasement < 0 {
						firstBasement = index + 1
					}
				}
			}
		}
	}
	fmt.Println(floor)
	fmt.Println(firstBasement)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
