package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./dec-02/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		total += processMatchResults(scanner.Text())
	}

	fmt.Println(total)
}

const (
	loss = 0
	draw = 3
	win  = 6
)

const (
	rock = iota + 1
	paper
	scissors
)

var wins = map[string]int{
	"A": paper, // paper beats rock. e.g. 2 points
	"B": scissors,
	"C": rock,
}

var loses = map[string]int{
	"A": scissors, // scissors loses to rock. e.g. 3 points
	"B": rock,
	"C": paper,
}

var draws = map[string]int{
	"A": rock, // scissors draws with scissors. e.g. 3 points
	"B": paper,
	"C": scissors,
}

func processMatchResults(result string) int {
	score := 0
	match := strings.Split(result, " ")
	opponent, self := match[0], match[1]

	switch self {
	case "X":
		score += loses[opponent]
		score += loss
	case "Y":
		score += draws[opponent]
		score += draw
	case "Z":
		score += wins[opponent]
		score += win
	}

	return score
}
