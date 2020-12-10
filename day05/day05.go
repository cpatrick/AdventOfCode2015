package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func hasThreeVowels(line string) bool {
	vowels := map[rune]struct{}{
		'a': struct{}{},
		'e': struct{}{},
		'i': struct{}{},
		'o': struct{}{},
		'u': struct{}{},
	}
	count := 0
	for _, c := range line {
		_, ok := vowels[c]
		if ok {
			count++
		}
		if count >= 3 {
			return true
		}
	}
	return false
}

func hasDoubleLetter(line string) bool {
	prevLetter := ' '
	for _, c := range line {
		if c == prevLetter {
			return true
		}
		prevLetter = c
	}
	return false
}

func noForbiddenStrings(line string) bool {
	re, err := regexp.Compile(".*(ab|cd|pq|xy).*")
	if err != nil {
		log.Fatal(err)
	}
	return !re.MatchString(line)
}

func isNice(line string) bool {
	return hasThreeVowels(line) && hasDoubleLetter(line) && noForbiddenStrings(line)
}

func doubles(line string) bool {
	found := make(map[string]int)
	prevLetter := ' '
	for i, c := range line {
		key := string(prevLetter) + string(c)
		val, ok := found[key]
		if ok {
			if i >= val+2 {
				return true
			}
		} else {
			found[key] = i
		}
		prevLetter = c
	}
	return false
}

func tweens(line string) bool {
	for i := 2; i < len(line); i++ {
		if line[i-2] == line[i] {
			return true
		}
	}
	return false
}

func isNiceNew(line string) bool {
	return doubles(line) && tweens(line)
}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var input []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	count := 0
	for _, line := range input {
		if isNice(line) {
			count++
		}
	}
	fmt.Println(count)

	count = 0
	for _, line := range input {
		if isNiceNew(line) {
			count++
		}
	}
	fmt.Println(count)
}
