package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const (
	// lower 01 - 26
	lower = "abcdefghijklmnopqrstuvwxyz"
	// upper 33 - 52
	upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func generatePriorities(alphabets ...string) map[rune]int {
	priorities := make(map[rune]int)
	i := 1
	for _, alphabet := range alphabets {
		for _, letter := range alphabet {
			priorities[letter] = i
			i++
		}
	}
	return priorities
}

func computeRucksackPriority(sack string, prio map[rune]int) int {
	first, second := splitString(sack)

	var match rune
	for _, letter := range first {
		if strings.Contains(second, string(letter)) {
			match = letter
		}
	}
	return prio[match]
}

func splitString(str string) (string, string) {
	mid := len(str) / 2
	first := make([]string, mid)
	for i := 0; i < mid; i++ {
		first[i] = string(str[i])
	}

	second := make([]string, mid)
	for i := mid; i < len(str); i++ {
		second[i-mid] = string(str[i])
	}

	return sortString(strings.Join(first, "")), sortString(strings.Join(second, ""))
}

func sortString(str string) string {
	runes := strings.Split(str, "")
	sort.Strings(runes)
	return strings.Join(runes, "")
}

func main() {
	priorities := generatePriorities(lower, upper)

	file, err := os.Open("./dec-03/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		total += computeRucksackPriority(line, priorities)
	}
	fmt.Println(total)
}
